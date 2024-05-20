package start

const initialBotMessage = `
Hello!
I am an open bot for visualizing the found spreads in the HermesTrade project.
You can choose the spread profitability you are interested in and the bot in real mode will send you all the found spreads once in a while. 
You can unsubscribe from the mailing list at any time using the appropriate command.

Enjoy and make a profit!
`

const availableCommandsWithDescription = `
Now bot supports several commands:

- /send_all : Receive all found spreads with profitability at least 0.5%. WARNING: There may be a lot of messages.
- /send_profitability_1 : Receive spreads with profitability from 1 percent.
- /send_profitability_5 : Receive spreads with profitability from 5 percent.
- /send_profitability_20 : Receive spreads with profitability from 20 percent. WARNING: These spreads appear extremely rarely.
- /stop : Unsubscribe from receiving all new spreads.
`
