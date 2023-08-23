# Online Boutique

This directory contains the port of the Google Cloud's [`Online
Boutique`][boutique] demo application to Service Weaver.

```mermaid
%%{init: {"flowchart": {"defaultRenderer": "elk"}} }%%
graph TD
    %% Nodes.
    github.com/ServiceWeaver/onlineboutique/adservice/T(adservice.T)
    github.com/ServiceWeaver/onlineboutique/cartservice/T(cartservice.T)
    github.com/ServiceWeaver/onlineboutique/cartservice/cartCache(cartservice.cartCache)
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T(checkoutservice.T)
    github.com/ServiceWeaver/onlineboutique/currencyservice/T(currencyservice.T)
    github.com/ServiceWeaver/onlineboutique/emailservice/T(emailservice.T)
    github.com/ServiceWeaver/onlineboutique/paymentservice/T(paymentservice.T)
    github.com/ServiceWeaver/onlineboutique/productcatalogservice/T(productcatalogservice.T)
    github.com/ServiceWeaver/onlineboutique/recommendationservice/T(recommendationservice.T)
    github.com/ServiceWeaver/onlineboutique/shippingservice/T(shippingservice.T)
    github.com/ServiceWeaver/weaver/Main(weaver.Main)

    %% Edges.
    github.com/ServiceWeaver/onlineboutique/cartservice/T --> github.com/ServiceWeaver/onlineboutique/cartservice/cartCache
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T --> github.com/ServiceWeaver/onlineboutique/cartservice/T
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T --> github.com/ServiceWeaver/onlineboutique/currencyservice/T
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T --> github.com/ServiceWeaver/onlineboutique/emailservice/T
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T --> github.com/ServiceWeaver/onlineboutique/paymentservice/T
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T --> github.com/ServiceWeaver/onlineboutique/productcatalogservice/T
    github.com/ServiceWeaver/onlineboutique/checkoutservice/T --> github.com/ServiceWeaver/onlineboutique/shippingservice/T
    github.com/ServiceWeaver/onlineboutique/recommendationservice/T --> github.com/ServiceWeaver/onlineboutique/productcatalogservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/adservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/cartservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/checkoutservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/currencyservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/productcatalogservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/recommendationservice/T
    github.com/ServiceWeaver/weaver/Main --> github.com/ServiceWeaver/onlineboutique/shippingservice/T
```

Here are the changes made to the original application:

* All of the services that weren't written in `Go` have been ported to `Go`.
* All of the networking calls have been replaced with the corresponding
  Service Weaver calls.
* All of the logging/tracing/monitoring calls have been replaced with the
  corresponding Service Weaver calls.
* The code is organized as a single Go module.

[boutique]: https://github.com/GoogleCloudPlatform/microservices-demo
