## newrelic entity tags create

Create tag:value pairs for the given entity

### Synopsis

Create tag:value pairs for the given entity

The create command adds tag:value pairs to the given entity.


```
newrelic entity tags create [flags]
```

### Examples

```
newrelic entity tags create --guid <entityGUID> --tag tag1:value1
```

### Options

```
  -g, --guid string   the entity GUID to create tag values on
  -h, --help          help for create
  -t, --tag strings   the tag names to add to the entity
```

### SEE ALSO

* [newrelic entity tags](newrelic_entity_tags.md)	 - Manage tags on New Relic entities

###### Auto generated by spf13/cobra on 8-Apr-2020