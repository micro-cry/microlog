# microlog

клиент-серверная система логирования с вшитой возможностью синхронизироватся с сервером и скидывать обновления логов
основной упор на оптимизацию места хранения

## [log_data] карта данных

| uid      | text   |
|----------|--------|
| [16]byte | string |
| {hash}   | -      |

## [log_instances] карта точек логирования

| uid      | text   |
|----------|--------|
| [16]byte | string |
| {hash}   | -      |

## [log_startop] старты

| uid      | start     | stop      | instance          | info         |
|----------|-----------|-----------|-------------------|--------------|
| [16]byte | time.Time | time.Time | [16]byte          | [16]byte     |
| {hash}   | -         | -         | log_instances.uid | log_data.uid |

## [log_status] карта статусов

| uid     | text   |
|---------|--------|
| [4]byte | string |
| {hash}  | -      |

## [log_path] карта путей

| uid      | text   |
|----------|--------|
| [12]byte | string |
| {hash}   | -      |

## [log_timeline]

| uid      | time      | status        | path         | stream          | data         |
|----------|-----------|---------------|--------------|-----------------|--------------|
| [32]byte | time.Time | [4]byte       | [8]byte      | [16]byte        | [16]byte     |
| {hash}   | -         | log_status.id | log_path.uid | log_startop.uid | log_data.uid |