# Open311-to-Go


## What is Open311?

Open311 is a form of technology that provides open channels of communication for issues that concern public space and public services. Primarily, Open311 refers to a standardized protocol for location-based collaborative issue-tracking. By offering free web API access to an existing 311 service, Open311 is an evolution of the phone-based 311 systems that many cities in North America offer.

Unlike the synchronous one-to-one communication of a 311 call center, Open311 technologies use the internet to enable these interactions to be asynchronous and many-to-many. This means that several different people can openly exchange information centered around a single public issue. This open model allows people to provide more actionable information for those who need it most and it encourages the public to be engaged with civic issues because they know their voices are being heard. Yet Open311 isn’t just about this more open internet-enabled model for 311 services, it’s also about making sure the technology itself is open so that 311 services and applications are interoperable and can be used everywhere.

Source: https://www.open311.org/learn/

## Why Go?

Most of the Open311 implementations so far are based on Python on Django framework. While easy to implement, that might not be the most optimal starting point for a highly performant and scalable API backend.

## Anything new here?

This implementation inludes the additions the City of Helsinki added on Open311, mainly support for other languages and locales and support for external media server in cases where images are included in the ticket.

Source: https://dev.hel.fi/apis/open311 

This implementation also uses MongoDB as a backend, also utilizing its spatial functions.

## Development framework and versions

* This work uses golang version 1.22.2. Due to recent development on the Go http/router package the latest version available is recommended.
* The API will be deployed as an Azure Function because it will then be easier to transfer to production platform.
* The development is done using Visual Studio Code - however it shouldn't make any difference what editor to use
* Sentry will be used for observations.

## What is the motivation for this?

This work is to support my master's thesis work on large scale asset management on urban digital twins. This should be up and running by the end of 2024.
