# covid-19-japan-api-cli

refer : [COVID-19 Japan Web API](https://documenter.getpostman.com/view/9215231/SzYaWe6h?version=latest)

# Use

## total command

```
go run main.go total
```

## prefectures command

```
go run main.go prefectures | sort -r -n -k 8 -t :
```
