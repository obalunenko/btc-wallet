# Backend Developer Test Assignment
Create a backend application that will allow users to register, create own BTC wallets and
transfer BTC to other wallets inside the platform. The platform makes 1.5% profit from the
transactions between users.

## The following RESTful API endpoints should be implemented:
● **POST /users** : creates user.

○ Returns a token that will authenticate all other requests for this user.

● **POST /wallets** : create BTC wallet for the authenticated user. 1 BTC (or 100000000
satoshi ) is automatically granted to the new wallet upon creation. User may register up to
10 wallets.

○ Returns wallet address and current balance in BTC and USD.


● **GET /wallets/:address** : returns wallet address and current balance in BTC and USD.


● **POST /transactions** : makes a transaction from one wallet to another

○ Transaction is free if transferred to own wallet.

○ Transaction costs 1.5% of the transferred amount (profit of the platform) if
transferred to a wallet of another user.

● **GET /transactions** : returns user’s transactions

● **GET /wallets/:address/transactions** : returns transactions related to the wallet


### Technical requirements:
● GO1.13+. Any open-source frameworks or libraries may be used.

● PostgreSQL or MySQL.

● BTC<->USD conversion should be implemented using any appropriate provider of your
preference. Rates should be up-to-date.

● Structure of the requests and responses should be decided by the developer.

● Implement authentication/authorization as simple as possible.

● Only API endpoints should be implemented, no frontend.
