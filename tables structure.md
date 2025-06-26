# The overall structure of the tables

This file is generated automatically

---

## data

| uid       | text   |
|-----------|--------|
| [16]bytes | string |
| _primary_ | -      |

## instances

| uid       | text   |
|-----------|--------|
| [16]bytes | string |
| _primary_ | -      |

## paths

| uid       | text   |
|-----------|--------|
| [12]bytes | string |
| _primary_ | -      |

## status

| uid       | text   |
|-----------|--------|
| [4]bytes  | string |
| _primary_ | -      |

## stream

| uid       | start    | stop     | instance      | info     |
|-----------|----------|----------|---------------|----------|
| [16]bytes | datetime | datetime | instances.uid | data.uid |
| _primary_ | _index_  | -        | _*index_      | _*index_ |

## timeline

| uid       | time     | stream     | status     | path      | data     |
|-----------|----------|------------|------------|-----------|----------|
| [18]bytes | datetime | stream.uid | status.uid | paths.uid | data.uid |
| _primary_ | -        | _*index_   | _*index_   | _*index_  | _*index_ |


