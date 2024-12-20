DROP TABLE IF EXISTS ACCOUNT;
DROP TABLE IF EXISTS SIMPLE_V2_ARB;
DROP TABLE IF EXISTS PAIR;
DROP TABLE IF EXISTS ERC20;
DROP TABLE IF EXISTS DEX;
DROP TABLE IF EXISTS NETWORK;

CREATE TABLE ACCOUNT(
    ACCOUNT_ID SERIAL PRIMARY KEY,
    NAME VARCHAR(50) NOT NULL,
    DESCRIPTION VARCHAR(150) NOT NULL,
    PUBLIC_KEY VARCHAR(100) NOT NULL,
    PRIVATE_KEY VARCHAR(150) NOT NULL,
    UNIQUE(ACCOUNT_ID)
);

INSERT INTO ACCOUNT( NAME, DESCRIPTION, PUBLIC_KEY, PRIVATE_KEY) VALUES
    ('TEST1', 'TEST DESC', '0xde6703A0F864487b04ac7BbaeD2B97381CB30050', '0xed2ae4cdb82cbe48af3f19b15d2ea9601115e6c090fe69f88897e252f30716a4'),
    ('Bot1', 'TEST DESC', '0x095708d3761CB68F6AAE5845c338263053214B27', 'cb4e4f7f4a14133b65693cac9acea58134e1324c9faee7b8d2f5051e994aa368'),
    ('Bot2', 'TEST DESC', '0x3626C815b5D30F0D8BDd2E12352f647a5d38C9a2', 'b3be17929896f80b8aa761e8b07ac25d01aaf80d3eee24d625de7e102f4a0c20'),
    ('Bot3', 'TEST DESC', '0x04199dfe8bFbdDa610022a659C3D4C61656B0D64', '06f896a4a05db4dc126fb2033e6c06ead6e1c65f21ffd54610ee095384ea699c'),
    ('Bot4', 'TEST DESC', '0x753E3427cf2a8FD4C8e1C701cBaAEAaedC7224DA', '8de1090a0169442d4bcb2fc35a59ef3800bf56f5beb89cbaa02ef912400125aa'),
    ('Bot5', 'TEST DESC', '0x87a34B99bBef08E206c54c01aD46BE2b75b7E89D', '077efb4563f92285cc2f1e628f1ab1fe7a8c78734da12c1dc8a6b5c798a255fc'),
    ('Bot6', 'TEST DESC', '0x1b0aa1A4AFd1a0C5a2C021db5e26bf79868aa783', '9286d4a2bee0527876dff53c3bf294ee90f8d0b4ac8438f7a1557590fa6c186f'),
    ('Bot7', 'TEST DESC', '0x0CC17504a91173891Acb16cdE55E99c1395ab439', '9ab9349a3e6dbabbf0f84d037bd5fb2d2d7a50a883da4149ef67b0b201891084'),
    ('Bot8', 'TEST DESC', '0x8eccAEcC4aCB031615F1b325dE62505067c5570D', '21216c6786059f6b45048de3b7c83b5710351f8d81229c9228c63589f7f18360'),
    ('Bot9', 'TEST DESC', '0x33f68a564Fe51e9ae215771ECCBcc92e2CE483fc', 'efaf353e06d1ce271b614f0967eb806cf6372a9ca966b95abe25ca1b6a04a60f');



CREATE TABLE NETWORK (
     NETWORK_ID SERIAL PRIMARY KEY,
     NAME VARCHAR(50) NOT NULL,
     CHAIN_ID INT NOT NULL,
     CURRENCY_SYMBOL VARCHAR(10) NOT NULL,
     BLOCK_EXPLORER_URL VARCHAR(150) NOT NULL,
     BASE_TOKEN_CONTRACT_ADDRESS VARCHAR(100) NOT NULL,
     UNIQUE(NETWORK_ID)
);

INSERT INTO NETWORK(NETWORK_ID, NAME, CHAIN_ID, CURRENCY_SYMBOL, BLOCK_EXPLORER_URL, BASE_TOKEN_CONTRACT_ADDRESS) VALUES
    (1, 'PulseChain_Network', 369, 'tPLS', 'https://rpc.pulsechain.com', '0xa1077a294dde1b09bb078844df40758a5d0f9a27'),
    -- ETHEREUM
    (10, 'Ethereum_Mainnet', 1, 'ETH', 'https://etherscan.io', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'),

    (11, 'Kovan_Test_Network', 42, 'ETH', 'https://kovan.etherscan.io', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'),
    (12, 'Goerli_Test_Network', 5, 'ETH', 'https://goerli.etherscan.io', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'),
    (13, 'Ropsten_Test_Network', 3, 'ETH', 'https://ropsten.etherscan.io', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'),
    (14, 'Rinkeby_Test_Network', 4, 'ETH', 'https://rinkeby.etherscan.io', '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2'),

    (40, 'BNB_CHAIN_mainnet', 56, 'BNB', 'https://bscnode1.anyswap.exchange', '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c'),


    (100, 'MATIC_MAINNET', 137, 'MATIC', 'https://explorer.matic.network', '');


CREATE TABLE ERC20 (
   ERC20_ID SERIAL PRIMARY KEY,
   NETWORK_ID SMALLINT NOT NULL,
   CONTRACT_ADDRESS VARCHAR(80) UNIQUE NOT NULL,
   NAME VARCHAR(150) NOT NULL,
   SYMBOL VARCHAR(150) NOT NULL,
   DECIMAL SMALLINT NOT NULL,
   SHOULD_FIND_ARB BOOLEAN DEFAULT FALSE NOT NULL,
   IS_VALIDATED BOOLEAN DEFAULT FALSE NOT NULL,
   CONSTRAINT FK_NETWORK
       FOREIGN KEY(NETWORK_ID)
           REFERENCES NETWORK(NETWORK_ID),
   UNIQUE(ERC20_ID)
);

CREATE TABLE DEX (
                     DEX_ID SERIAL PRIMARY KEY,
                     NAME VARCHAR(50) NOT NULL UNIQUE,
                     NETWORK_ID SMALLINT NOT NULL,
                     ROUTER_ADDRESS VARCHAR(100) NOT NULL,
                     FACTORY_ADDRESS VARCHAR(100) NOT NULL,
                     FACTORY_ABI TEXT,
                     ROUTER_ABI TEXT,
                     CONSTRAINT FK_NETWORK
                         FOREIGN KEY(NETWORK_ID)
                             REFERENCES NETWORK(NETWORK_ID),
                     UNIQUE(DEX_ID)
);

INSERT INTO DEX(DEX_ID, NETWORK_ID, NAME, ROUTER_ADDRESS, FACTORY_ADDRESS) VALUES
      (1, 1, 'PulseChain_Network_Uniswap_V2', '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D', '0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f'),
      (2, 1, 'PulseChain_Network_SushiSwap_V2', '0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F', '0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac'),
      (3, 1, 'PulseChain_Network_PulseX_V1', '0x98bf93ebf5c380C0e6Ae8e192A7e2AE08edAcc02', '0x1715a3E4A142d8b698131108995174F37aEBA10D'),
      (4, 1, 'PulseChain_Network_PulseX_V2', '0x165C3410fC91EF562C50559f7d2289fEbed552d9', '0x29ea7545def87022badc76323f373ea1e707c523'),
      (5, 1, 'PulseChain_ShibaSwap', '0x03f7724180AA6b939894B5Ca4314783B0b36b329', '0x115934131916C8b277DD010Ee02de363c09d037c'),
      (6, 1, '9mm_v2', '0xcC73b59F8D7b7c532703bDfea2808a28a488cF47', '0x3a0Fa7884dD93f3cd234bBE2A0958Ef04b05E13b'),
      (7, 1, '9inch_v2', '0xeB45a3c4aedd0F47F345fB4c8A1802BB5740d725', '0x5b9F077A77db37F3Be0A5b5d31BAeff4bc5C0bD7'),
      (8, 1, 'SparkSwap', '0x76C08825b4A675FD6a17A244660BabeB4ADA79d5', '0x955219A87eB0C6754fd247266af970F7d16906CD'),
      (9, 1, 'PulseRate', '0x71bb8a2feD36aa2dEa9f8f9Cb43E038315Dd7ba3', '0xE2332E5297b18a21DcE0E6ac461e821C353A00cA'),
      (10, 1, 'Ezswap', '0x05d5F20500eD8d9E012647E6CFe1b2Bf89f5b926', '0xbeE82c502eC48a962510Aa4505595259C3ba631f'),


      (20, 10, 'Ethereum_Mainnet_Uniswap_V2', '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D', '0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f'),
      (21, 10, 'Ethereum_Mainnet_SushiSwap_V2', '0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F', '0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac'),

      (30, 11, 'Kovan_Test_Network_Uniswap_V2', '0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D', '0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f'),


      (50, 40, 'BNB_CHAIN_mainnet_PancakeSwap_V2',  '0x10ED43C718714eb63d5aA57B78B54704E256024E', '0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73'),
      (51, 40, 'BNB_CHAIN_mainnet_BiSwap_V2',  '0x3a6d8cA21D1CF76F653A67577FA0D27453350dD8', '0x858E3312ed3A876947EA49d572A7C42DE08af7EE'),
      (52, 40, 'BNB_CHAIN_mainnet_ApeSwap_V2',  '0xcF0feBd3f17CEf5b47b0cD257aCf6025c5BFf3b7', '0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6'),

      (100, 100, 'MATIC_QUICKSWAP_V2',  '0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff', '0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32');

CREATE TABLE PAIR(
    PAIR_ID         SERIAL               PRIMARY KEY,
    DEX_ID          INTEGER               NOT NULL,
    PAIR_INDEX          INTEGER               NOT NULL,
    TOKEN0_ID       INTEGER               NOT NULL,
    TOKEN1_ID       INTEGER               NOT NULL,
    PAIR_CONTRACT_ADDRESS VARCHAR(100) NOT NULL,
    TOKEN0_RESERVES NUMERIC(78, 0),
    TOKEN1_RESERVES NUMERIC(78, 0),
    SHOULD_FIND_ARB BOOLEAN DEFAULT FALSE NOT NULL,
    UNIQUE(DEX_ID, TOKEN0_ID, TOKEN1_ID),
    CONSTRAINT FK_DEX
        FOREIGN KEY (DEX_ID)
            REFERENCES DEX (DEX_ID),
    CONSTRAINT FK_ERC20_0
        FOREIGN KEY (TOKEN0_ID)
            REFERENCES ERC20 (ERC20_ID),
    CONSTRAINT FK_ERC20_1
        FOREIGN KEY (TOKEN1_ID)
            REFERENCES ERC20 (ERC20_ID),
    UNIQUE (PAIR_ID)
);


CREATE TABLE SIMPLE_V2_ARB(
     SIMPLE_V2_ARB_ID         SERIAL               PRIMARY KEY,
     PAIR0_ID       INTEGER               NOT NULL,
     PAIR1_ID       INTEGER               NOT NULL,
     SHOULD_FIND_ARB BOOLEAN DEFAULT FALSE NOT NULL,
     FAILED_ARB_COUNT INTEGER DEFAULT 0,
     ERROR_ARB_COUNT INTEGER DEFAULT 0,
     NOTE VARCHAR(100),
     UNIQUE(PAIR0_ID, PAIR1_ID),
     CONSTRAINT FK_SIMPLE_V2_ARB_PAIR0
         FOREIGN KEY (PAIR0_ID)
             REFERENCES PAIR (PAIR_ID),
     CONSTRAINT FK_SIMPLE_V2_ARB_PAIR1
         FOREIGN KEY (PAIR0_ID)
             REFERENCES PAIR (PAIR_ID),
     UNIQUE (SIMPLE_V2_ARB_ID)
);
CREATE UNIQUE INDEX SIMPLE_V2_ARB_PAIR_ID_idx ON SIMPLE_V2_ARB (PAIR0_ID, PAIR1_ID);


CREATE TABLE TRIANGLE_ARB(
     TRIANGLE_ARB_ID  SERIAL           PRIMARY KEY,
     PAIR0_ID            INTEGER          NOT NULL,
     PAIR1_ID            INTEGER          NOT NULL,
     PAIR2_ID            INTEGER          NOT NULL,
     SHOULD_FIND_ARB BOOLEAN DEFAULT FALSE NOT NULL,
     FAILED_ARB_COUNT INTEGER DEFAULT 0,
     SUCCESS_ARB_COUNT INTEGER DEFAULT 0,
     NOTE VARCHAR(100),
     UNIQUE(PAIR0_ID, PAIR1_ID, PAIR2_ID),
     CONSTRAINT FK_TRIANGLE_V2_ARB_PAIR0
         FOREIGN KEY (PAIR0_ID)
             REFERENCES PAIR (PAIR_ID),
     CONSTRAINT FK_TRIANGLE_V2_ARB_PAIR1
         FOREIGN KEY (PAIR1_ID)
             REFERENCES PAIR (PAIR_ID),
     CONSTRAINT FK_TRIANGLE_V2_ARB_PAIR2
         FOREIGN KEY (PAIR2_ID)
             REFERENCES PAIR (PAIR_ID),
     UNIQUE (TRIANGLE_ARB_ID)
);


