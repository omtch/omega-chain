# Omega智能链Omega Chain

简体中文 | [English](README.md)

**Omega智能链介绍**

Omega Chain是一条去中心化高效节能公链，可为开发人员提供高效且低成本的链上环境，以运行去中心化智能合约应用程序(DApps)和存储数字资产。

Omega Chain是Omega立足区块链生态资源，基于Omega技术和新项目挖掘优势，构建的开放友好区块链平台。Omega Chain 生态联盟为多链Dapp提供全方位支持服务，帮助开发者更好服务用户需求，提供低成本与高效能并存的优质链上体验。

Omega Chain采用PoSA共识机制，每个参与节点拥有相同的权利，它可以让开发者自由构建去中心化应用，包括DeFi产品、DAPP、数字资产等。此外通过跨链模块，可以简单高效地实现各区块链的价值互联互通，从而共同构建生态系统和增值系统。

Omega Chain旨在促进基于区块链技术的大规模商业应用程序的开发，Omega Chain生态联盟共建开放共生生态，致力于拓宽区块链世界并带来更加完整的生态闭环。



**Omega Smart Chain（Omega Chain）设计逻辑**

Omega智能链采用高度抽象的模块化设计思路，将系统划分为基础网络、数据库存储、共识算法、交易处理机、虚拟机、应用层接口等几个核心模块。

Omega Chain公链上线后，任何开发者都可以便捷的在Omega Chain链上构建去中心化项目，另外Omega Chain后期会逐步开放若干创新性的服务。


**设计逻辑**

1、从技术上来讲，Omega Chain不属于二层解决方案，是一个独立的公链。大多数Omega Chain的技术功能以及业务功能由Omega Chain团队开发。

2、可兼容包括以太坊、BSC、HECO等当前主流公链。以上公链具有相对成熟的应用以及社区。因此行业内的大多数成熟的DAPP、生态系统组件和工具可以和Omega Chain相适配。Omega Chain节点将进行硬件规范以及Omega Chain的功能运行。

****

**改进方案**

根据当前行业面临的问题，Omega Chain将在以下几个层面进行改善

1、区块确认是3秒，这将比包括以太坊、BSC在内的主流区块链都快

2、Omega Chain的手续费用充当区块奖励，手续费用由OMN进行支付，此举拓宽OMN的价值。

3、吸收与当前主流公链的优势并与其充分竞争

4、引入Staking机制以及Omega Chain治理机制



**Omega Smart Chain（Omega Chain）特点**

共识机制：PoSA

TPS：500+

出块时间：3秒

**特点：**

1、高吞吐量

高吞吐量是通过改善Omega Chain中的TPS实现的，Omega Chain区块确认时间为3秒，日常使用实用程度已经超过比特币和以太坊。

2、可扩展性

基于良好的可扩展性和高效的智能合约，应用程序可以在Omega Chain中有更多部署方式，Omega Chain可以支持大量的用户。

3、高可靠性

Omega Chain具有更可靠的网络结构，用户资产，内在价值，并且更高程度的分权化共识带来了改进的奖励分配机制。



**PoSA**

共识机制种类大不相同，包括Proof-of-Work（POW）、Proof-of-Authority（POA）、Proof-of-Stake（POS）。POW通过算力进行挖矿来维护网络，POA采用验证人机制，不过这被一部分人认为POA去中心化的程度较低。Omega Chain将采用的PoSA共识机制，此机制融合了POA、POS的特点，其中：

1、验证者数量有限，区块由一定数量的验证者验证后产出

2、验证者轮流产生区块，类似于POA的产生方式

3、可以通过Staking成为验证者参与Omega Chain的治理



**跨链技术**

跨链技术原理的利用对于区块链行业至关重要，通过Omega Chain跨链技术可以使用户资产自由流通，优势包括：

1、用户可基于Omega Chain搭建数字资产、去中心化金融产品

2、Omega Chain链上项目、资产可自由稳定地流通、并且比当前行业内的各大公链更加高效、便捷、以及成本低廉。

3、Omega Chain可承担区块链资产中转站的角色，通过资产跨链桥，将各公链资产映射到Omega Chain，在链上锁定资产后，在Omega Chain生成对应数量的Token



**Omega虚拟机（OVM）**

Omega智能链实现的虚拟机全面兼容以太坊虚拟机，方便开发者移植现有的DAPP，不仅降低了开发者的学习成本，并且由于Omega智能链PoSA共识算法天然的优势，大幅提高了DAPP的运行效率，同时运行成本大幅降低。

Omega智能链的虚拟机还进行了诸多优化，使得DAPP的运营成本大幅降低，同时开发了许多新特性来支持智能合约的业务逻辑，比如在智能合约中支持批量验签、在智能合约中支持合约地址的判断等。

1、轻量级

OVM采用轻量级的虚拟机构架，旨在节省运行空间，减少资源耗费及保证系统性能。

2、稳定、安全性

OVM采用了严谨的设计规范，低粒度的底层操作码，保证了每个计算步骤的精确性，最大程度消除产生歧义的空间。 同时出于安全性的考量，OVM的转账与运行合约均不需要消耗代币，只会消耗带宽，避免了针对类似以太坊gas消耗模式的攻击。在保证了每个操作计算步骤的确定性的同时，也保证了带宽消耗的稳定性。

3、兼容性

目前，OVM能完美兼容以太坊EVM，并在未来兼容更多主流的VM。因此, 以太坊上的智能合约，能直接运行到OVM上，无缝对接现有的开发者生态，提高开发者的开发效率。开发者无需学习新的编程语言，就能用Solidity等主流编程语言在熟悉的Remix环境中进行智能合约的开发、调试、编译，将极大缩减开发成本。

4、开发人员友好性

OVM的带宽消耗模式减少了合约的开发成本。让开发者可以把更多精力放在合约代码的逻辑本身。同时，OVM提供了对开发者友好的一站式的部署、触发、查看智能合约的接口。



**模型设计与机制**

**账户模型**

Omega Chain采用账户模型。账户的唯一标识为地址address，对账户操作需要验证私钥签名。项目方可以发布并创建智能合约，也可以调用他人发布的智能合约，也可以对节点进行投票等等。Omega Chain所有的活动都围绕账户进行。



**资源模型**

Omega智能链设计了一套完善的资源模型，并支持资源模型参数的动态调整，这是一个非常优秀的反馈调节机制，比如当链上交易繁忙的时候，交易手续费用使用费用会变的较高，当空闲时，这些资源的成本会随之降低。另外还为用户设置一定量的免费资源配额。

同时支持用户通过冻结OMN的方式来获取相应数量的投票权，用户通过为节点投票可以获得相应的奖励。



****

**链上治理和投票机制**

Omega智能链设置了科学高效的激励机制，促进区块链的自我繁荣，节点有权利生产区块，打包交易，并获取相应的区块生产激励，同时节点还可以获得选票奖励。



系统参数可以通过社区的治理进行调控，包括：

1、Omega Chain系统合约的参数都是灵活的，如跨链转移费用，中继器奖励金额等，这有利于生态的良好运行

2、Omega Chain上的Stake / Slash / Oracle模块的参数

所有这些参数由Omega Chain验证程序集根据其Staking通过提议投票过程一起确定，这样的过程都将在Omega Chain链上进行。



治理设计原则

1、统一的界面，对这些参数感兴趣的合同仅需要实现相同的接口。

2、可扩展，添加新的系统合同时，无需修改任何其他合同。

3、失败容忍度，验证者可以投票跳过错误的建议并继续。

4、多路复用，现在我们使用参数gov，但是将来会有更多的治理功能。



投票流程：

1、选票

Omega Chain中设定持有OMN就可以拥有选票的权利。

2、投票过程

Omega Chain设定对候选人的投票是一笔特殊类型的交易，节点可以通过生成一笔投票交易对候选人进行投票。

3、统计票

每个维护期内，统计一次候选人的票数，将获得票数最多的候选人作为下一个出块周期的记账人。



**激励机制**

为了保证区块链系统安全高效地运行，Omega智能链设定激励模型用于鼓励更多的节点加入到Omega Chain网络中，从而扩大网络规模，对于记账人当他们完成出块任务，给予相应的OMN奖励。Omega Chain设定witness每生产一个被固化的区块，就会获得一定的OMN奖励；对于所有记账人，每个Epoch的维护期会依据得票率的多少分配固定的奖励。并且激励的数量是透明的，激励的发放过程是完全去中心化的。


**Omega Chain Token标准**

Omega Chain token是通过Omega Chain链上发行的标准代币。Omega Chain完全兼容以太坊Ethereum、BSC、HECO标准，因此Omega Chain链同样可支持ERC20 Token、BSC Token、Heco Token。用户可在Omega交易所内部或者Omega Chain跨链桥进行链间互换。



**Omega Chain生态**

Omega智能链是一个规模庞大的分布式操作系统，全世界范围内有成千上万的节点平稳地运行在遍布全球的服务器和终端之上。Omega智能链的强大离不开应用软件的支持，Omega智能链生态的繁荣离不开社区开发者的追随。Omega生态系统中拥有包括DeFi、DAPP、NFT等板块在内的众多优秀的应用。如钱包、区块链浏览器、DEX、借贷、预言机、NFT交易市场等。


**共识机制**

Omega Chain采用PoSA共识机制，具有交易成本低、交易延时低、交易并发高等特点，支持最多21个验证人节点；

**名词解释**

验证人：负责对链上交易进行打包出块；

活跃验证人：即当前负责打包出块的一组验证人，上限为21个。

Epoch：以区块为单位的时间间隔，目前Omega Chain上 1epoch = 200block，每个epoch结束的时候，区块链会与系统合约交互，进行活跃验证人更新；

****

**系统合约**

目前验证人的管理，均由系统合约完成，目前的系统合约有：

Proposal 负责管理验证人的准入资格，管理验证人提案和投票；

Validators 负责对验证人进行排名管理、质押和解质押操作、分发区块奖励等；

Punish 负责对不正常工作的活跃验证人进行惩罚操作；

区块链调用系统合约：

每个块结束的时候，会调用Validators合约，将区块中所有交易的手续费分发给active validator;

发现validator没有正常工作的时候，会调用Punish合约，对validator进行惩罚；

每个epoch结束的时候，会调用Validators合约，根据排名，更新active validator；


**质押**

任何账户，都可以对validator进行任意数量的质押操作，每个validator的最小质押量是100000OMN。 如果想取回已质押的OMN，需要按照如下操作进行：

发送调用Validators合约，发送针对某一个validator的解质押(unstake)的声明交易;

等待86400个块之后，调用Validators合约，发送提取质押(withdrawStaking)的交易，将所有在此validator的质押取回；


**惩罚措施**

每当发现验证人没有按照预先设定进行出块的时候，就会在这个块结束时，自动调用Punish合约，对验证人进行计数。当计数达到24时，罚没验证人的所有收入。当计数达到48时，将验证人移除出活跃验证人列表，同时取消验证人资格。

