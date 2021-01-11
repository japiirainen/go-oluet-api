# Oluet-api

API for finding information about you favorive alcoholic beverages!😎
(not only beer!)

![cool beer picture](https://cdn1.iconfinder.com/data/icons/zaficons-foods-1/512/beer-512.png) 

## This is a rewrite of my oluet-api in go!

- The poject was originally written in nodejs, but I decided to start learning glo and thought this was a good place to start exploring!

## Links
- [Homepage](https://oluet-api.xyz)
- [Graphql playground](https://oluet-api.xyz/graphql)
- [API](https://oluet-api.xyz/query)

If you don't know graphql playground is a great place to try out some queries straight from you favorite browser. Give it a shot!

## API documentation v1

### Graphql endpoint

This endpoint is used by all of the queries. You should read the
official graphql documentation on more information about the
technology if it's new for you!
[Graphal docs](https://graphql.org)
```
https://oluet-api.xyz/query
```

### Graphql playground endpoint

If you are not yet familiar with graphql you are going to love the
playground! There is no need to use tools like Postman with graphql
since you can test your queries straight from the playground! You
should also see the docs for playground to utilize it's full
potential!
[Graphal playground docs](https://github.com/graphql/graphql-playground)
```
https://oluet-api.xyz/graphql
```

### Sample queries

You can experiment with these at the playground!

### List all drinks
```
query {
    drinks {
        nimi
        date
        productID
        hinta
        tyyppi
    }
}
```

### List all beers
```
query {
    beers {
        nimi
        date
        olutTyyppi
        hinta
        tyyppi
    }
}
```

### Search for all "karhu" beers
```
query {
    beersearch(term: "karhu") {
        nimi
        date
        valmistaja
        huomautus
        olutTyyppi
        hinta
        tyyppi
    }
}
```

### Search for all "karhu" beers
```
query {
    pricehistory(productID:"792176") {
        hinta
        date 
    }
 }
```

## TODO
1. maybe make a separate type for olut