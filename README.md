# fx_decorate_value_group_issue

This repo contains the code to reproduce the issue [#1104](https://github.com/uber-go/fx/issues/1104) in [uber-go/fx](https://github.com/uber-go/fx).

## Reproduction

It's as simple as

```bash
go run main.go
```

## Expected output

I **expect** the logs to look sth like the following.
I reduced the output only to the relevant fields.

```
{"logger":"my_service.service","msg":"name should be `my_service.service`"}
{"logger":"my_service.dummy","msg":"name should be `my_service.dummy`"}
```

## Actual output

However, this is the **actual**, faulty output, again only the relevant fields:

```
{"logger":"my_service.service","msg":"name should be `my_service.service`"}
{"logger":"dummy","msg":"name should be `my_service.dummy`"}
```
