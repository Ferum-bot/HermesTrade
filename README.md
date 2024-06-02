# HermesTrade

There are many online exchanges in the world that allow you to conduct transactions and buy financial assets over the Internet. 
Different exchanges are focused on different regions and have different liquidity. 
In addition, the demand for the same assets on different exchanges at one time may be different, so the cost may vary significantly. 

So, The project “**HermesTrade**” is aimed at creating an automated system that will analyze the value of various financial assets on different exchanges in real time 
and look for a sequence of transactions with which you can extract due to the difference in supply and demand on different exchanges.


## Features:

### Unique search algorithm

This algorithm is located in a separate external package [pkg/asset-spread-hunter](./pkg/asset-spread-hunter).
It basically contains two algorithms with a lot of heuristics and meta-heuristics.

Full information and detailed description of the algorithm, including stress testing and load measurements are located in
package [README.md](pkg/asset-spread-hunter/README.md)

### High throughput

The average throughput of the project is about 150 thousand financial quotes per minute.

Thanks to the speed of work, it turns out to analyze a huge part of the data and find really profitable spreads.


### Support for 6 exchanges

Now totally supports six different finance exchanges:
* [Binance](https://www.binance.com/en)
* [ByBit](https://www.bybit.com/en/)
* [Coinbase](https://www.coinbase.com/)
* [Kraken](https://www.kraken.com/)
* [OKX](https://www.okx.com/en)
* [Upbit](https://upbit.com/home)

### Convenient and concise visualizer

Now project supports simple Telegram Bot for visualizing the found spreads. 
You can search our bot in Telegram by bot name: `@HermesTradeOpenBot  `
Or just click [here](https://t.me/HermesTradeOpenBot). The bot will tell you how to use it.

## Project Architecture:

The architecture of the project consists of several related services divided by areas of responsibility.

In total, the project presents 5 types of services:

- Scrappers - collect information of currency pairs cost from one specific financial exchange
- AssetsStorage - deals  with the unification and pre-storage of currency pairs data.
- AssetSpreadHunter - performs the basic business logic of the project and analyzes financial assets and searches profitable spreads
- SpreadsStorage - responsible for storing spreads found all the time and providing them to external services
- Connectors - project services that provides found spreads for external sources(for example telegram bot)

The architecture is more clearly demonstrated in the diagram below:

<img src='./images/HermesTrade_Architecture.png' width='724' height='524'>

## How to launch?

To run the project locally, you will need to install several programs:

* Docker and Docker-compose
* Golang
* MiniKube

After installation, you need to use the commands from the Makefile:

First of all, launch databases and message queues:

`make launch-dev-env`

After that, launch main services:

1. `make launch-assets-storage`
2. `make launch-spreads-storage`
3. `make launch-asset-spread-hunter`

Then launch scrappers:

`make launch-all-scrappers`

And in the end launch telegram connector:

`make launch-connector-telegram`


## How to contribute?

1. Create new branch from main: `git branch <YOUR_NICKNAME>:<FEATURE_NAME>`
2. Checkout to your branch: `git checkout <BRANCH_NAME_FROM_POINT_1>`
3. Write code
4. Write tests(__IMPORTANT__)
5. Test code on development environment
6. Create Pull Request
7. Wait for approve
