// build +unit

package recipes

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/newrelic/newrelic-cli/internal/install/types"
)

var (
	discoveryManifest types.DiscoveryManifest
	recipeCache       []types.OpenInstallationRecipe
	repository        *RecipeRepository
)

func Setup() {
	discoveryManifest = types.DiscoveryManifest{}
	recipeCache = []types.OpenInstallationRecipe{}
	repository = NewRecipeRepository(recipeLoader)
}

func Test_ShouldFindAll_Empty(t *testing.T) {
	Setup()
	recipes, _ := repository.FindAll(discoveryManifest)

	require.Empty(t, recipes)
}

func Test_ShouldFindSingleRecipe(t *testing.T) {
	Setup()
	givenCachedRecipe("id1", "my-recipe")

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 1)
	require.Equal(t, results[0].ID, "id1")
}

func Test_ShouldFindSingleOsRecipe(t *testing.T) {
	Setup()
	givenCachedRecipeOs("id1", "my-recipe", types.OpenInstallationOperatingSystemTypes.LINUX)
	discoveryManifest.OS = "linux"

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 1)
	require.Equal(t, results[0].ID, "id1")
}

func Test_ShouldNotFindSingleOsRecipe(t *testing.T) {
	// log.SetLevel(log.TraceLevel)
	Setup()
	givenCachedRecipeOs("id1", "my-recipe3", types.OpenInstallationOperatingSystemTypes.LINUX)

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 0)
}

func Test_ShouldFindMostMatchingSingleRecipe(t *testing.T) {
	Setup()
	givenCachedRecipeOs("id1", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX)
	givenCachedRecipeOsPlatform("id2", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.DEBIAN)
	givenCachedRecipeOsPlatform("id3", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.UBUNTU)
	discoveryManifest.OS = "linux"
	discoveryManifest.Platform = "debian"

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 1)
	require.Equal(t, results[0].ID, "id2")
}

func Test_ShouldFindMostMatchingSingleRecipeWithoutPlatform(t *testing.T) {
	Setup()
	givenCachedRecipeOs("id1", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX)
	givenCachedRecipeOsPlatform("id2", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.DEBIAN)
	givenCachedRecipeOsPlatform("id3", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.UBUNTU)
	discoveryManifest.OS = "linux"
	discoveryManifest.Platform = "centos"

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 1)
	require.Equal(t, results[0].ID, "id1")
}

func Test_ShouldDiscardMostMatchingWithoutAllFieldsMatching(t *testing.T) {
	Setup()
	givenCachedRecipeOs("id1", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX)
	givenCachedRecipeOsPlatform("id2", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.DEBIAN)
	givenCachedRecipeOsPlatform("id3", "my-recipe2", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.UBUNTU)
	givenCachedRecipeOsPlatformVersionArch("id4", "my-recipe2", types.OpenInstallationOperatingSystemTypes.WINDOWS, "10.0", "x86_64")
	discoveryManifest.OS = "linux"
	discoveryManifest.PlatformVersion = "10.0"
	discoveryManifest.KernelArch = "x86_64"

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 1)
	require.Equal(t, results[0].ID, "id1")
}

func Test_ShouldFindMultipleNames(t *testing.T) {
	Setup()
	givenCachedRecipeOs("id1", "infra", types.OpenInstallationOperatingSystemTypes.LINUX)
	givenCachedRecipeOsPlatform("id2", "infra", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.DEBIAN)
	givenCachedRecipeOsPlatform("id3", "logs", types.OpenInstallationOperatingSystemTypes.LINUX, types.OpenInstallationPlatformTypes.UBUNTU)
	givenCachedRecipeOsArch("id4", "logs", types.OpenInstallationOperatingSystemTypes.LINUX, "x86_64")
	discoveryManifest.OS = "linux"
	discoveryManifest.Platform = "debian"
	discoveryManifest.KernelArch = "x86_64"

	results, _ := repository.FindAll(discoveryManifest)

	require.Len(t, results, 2)
	require.True(t, containsId(results, "id2"))
	require.True(t, containsId(results, "id4"))
}

func Test_matchRecipeCriteria_Basic(t *testing.T) {
	Setup()
	discoveryManifest.Platform = "linux"

	hostMap := getHostMap(discoveryManifest)
	actual := matchRecipeCriteria(hostMap, "Platform", "linux")
	require.True(t, actual)
}

func Test_matchRecipeCriteria_EmptyString(t *testing.T) {
	Setup()
	hostMap := getHostMap(discoveryManifest)
	actual := matchRecipeCriteria(hostMap, "Platform", "")
	require.True(t, actual)
}

func Test_matchRecipeCriteria_KeyMissing(t *testing.T) {
	Setup()

	hostMap := getHostMap(discoveryManifest)
	actual := matchRecipeCriteria(hostMap, "KeyMissing", "xyz")
	require.False(t, actual)
}

func recipeLoader() ([]types.OpenInstallationRecipe, error) {
	return recipeCache, nil
}

func givenCachedRecipeOs(id string, name string, os types.OpenInstallationOperatingSystem) *types.OpenInstallationRecipe {
	r := createRecipe(id, name)
	t := types.OpenInstallationRecipeInstallTarget{
		Os: os,
	}
	r.InstallTargets = append(r.InstallTargets, t)
	recipeCache = append(recipeCache, *r)
	return r
}

func givenCachedRecipeOsPlatform(id string, name string, os types.OpenInstallationOperatingSystem, platform types.OpenInstallationPlatform) *types.OpenInstallationRecipe {
	r := createRecipe(id, name)
	t := types.OpenInstallationRecipeInstallTarget{
		Os:       os,
		Platform: platform,
	}
	r.InstallTargets = append(r.InstallTargets, t)
	recipeCache = append(recipeCache, *r)
	return r
}

func givenCachedRecipeOsArch(id string, name string, os types.OpenInstallationOperatingSystem, arch string) *types.OpenInstallationRecipe {
	r := createRecipe(id, name)
	t := types.OpenInstallationRecipeInstallTarget{
		KernelArch: arch,
		Os:         os,
	}
	r.InstallTargets = append(r.InstallTargets, t)
	recipeCache = append(recipeCache, *r)
	return r
}

func givenCachedRecipeOsPlatformVersionArch(id string, name string, os types.OpenInstallationOperatingSystem, platformVersion string, arch string) *types.OpenInstallationRecipe {
	r := createRecipe(id, name)
	t := types.OpenInstallationRecipeInstallTarget{
		KernelArch:      arch,
		Os:              os,
		PlatformVersion: platformVersion,
	}
	r.InstallTargets = append(r.InstallTargets, t)
	recipeCache = append(recipeCache, *r)
	return r
}

func givenCachedRecipe(id string, name string) *types.OpenInstallationRecipe {
	r := createRecipe(id, name)
	recipeCache = append(recipeCache, *r)
	return r
}

func createRecipe(id string, name string) *types.OpenInstallationRecipe {
	r := &types.OpenInstallationRecipe{
		ID:   id,
		Name: name,
	}
	return r
}

func containsId(recipes []types.OpenInstallationRecipe, id string) bool {
	for _, recipe := range recipes {
		if recipe.ID == id {
			return true
		}
	}

	return false
}
