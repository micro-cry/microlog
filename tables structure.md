# The overall structure of the tables
This file is generated automatically

---

## data

| uid       | text   |
|-----------|--------|
| [16]bytes | string |
| _primary_ | _-_    |


## instances

| uid       | text   |
|-----------|--------|
| [16]bytes | string |
| _primary_ | _-_    |


## paths

| uid       | text   |
|-----------|--------|
| [12]bytes | string |
| _primary_ | _-_    |


## status

| uid       | text   |
|-----------|--------|
| [4]bytes  | string |
| _primary_ | _-_    |


## stream

| uid       | start    | stop     | instance      | info     |
|-----------|----------|----------|---------------|----------|
| [16]bytes | datetime | datetime | instances.uid | data.uid |
| _primary_ | _-_      | _-_      | _index_       | _index_  |


## timeline

| uid       | time     | stream     | status     | path      | data     |
|-----------|----------|------------|------------|-----------|----------|
| [18]bytes | datetime | stream.uid | status.uid | paths.uid | data.uid |
| _primary_ | _-_      | _index_    | _index_    | _index_   | _index_  |


