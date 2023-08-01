# Examples

The following examples are provided to help users get started with dsl-parser.
They are arranged as follows:

* `dataset` - a simple example showing how to use dsl-parser to query database

# Usage
The dsl parser api presented as below
```go
    Parser.Parse(dataset, payload)
```
 where dataset contains the name of database schema and the database type, payload contains the statement of dsl
\
\
\
Let's look at how our parser generates sql

```go
    Dataset{
        Table: "table1",
        Type:  parser.DatasetTypeTable,
    }
    Payload{
        "workflow": [
            {
                "type": "view",
                "query": [
                        {
                            "op": "raw",
                            "fields": [
                                "*"
                            ]
                        }
                ]
            }
        ],
        "offset": 0,
        "limit": 100
    }
    // "SELECT * FROM table1 LIMIT 100 OFFSET 0"
    sql := baseParser.Parse(dataset, payload)
```