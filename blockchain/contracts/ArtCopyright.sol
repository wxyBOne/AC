// SPDX-License-Identifier:MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract ArtCopyright is ERC721 {
    // 艺术品结构
    struct Artwork {
        uint256 artworkId;      // 艺术品ID，用于唯一标识每个艺术品
        string title;           // 标题，存储艺术品的名称
        string description;     // 描述，用于详细介绍艺术品的相关情况
        uint256 price;         // 价格，记录艺术品的售卖价格
        address creator;        // 创作者，指向创建该艺术品的以太坊地址
        uint256 createTime;     // 创建时间，记录艺术品在区块链上被创建的时间戳（以秒为单位）
        bool isForSale;        // 是否在售，用于表示艺术品当前是否处于可售卖状态
    }

    // 存储所有艺术品
    mapping(uint256 => Artwork) private _artworks;

    // 事件声明
    event ArtworkRegistered(uint256 indexed artworkId, address indexed creator, uint256 price);
    event ArtworkSold(uint256 indexed artworkId, address indexed from, address indexed to, uint256 price);
    event PriceUpdated(uint256 indexed artworkId, uint256 oldPrice, uint256 newPrice);

    // 调用父合约ERC721的构造函数，传入合约名称"ArtCopyright"和符号"ART"，用于初始化ERC721相关的设置
    constructor() ERC721("ArtCopyright", "ART") {}

    // 注册新艺术品
    function registerArtwork(
        uint256 artworkId,
        string memory title,
        string memory description,
        uint256 price
    ) public returns (uint256) {
        // 检查artworkId是否已存在
        require(_artworks[artworkId].creator == address(0), "Artwork ID already exists");

        // 调用从ERC721合约继承来的_safeMint函数，将艺术品的所有权初始分配给调用者（msg.sender），并使用传入的artworkId作为唯一标识进行铸造
        _safeMint(msg.sender, artworkId);

        _artworks[artworkId] = Artwork({
            artworkId: artworkId,
            title: title,
            description: description,
            price: price,
            creator: msg.sender,
            createTime: block.timestamp,
            isForSale: true
        });

        emit ArtworkRegistered(artworkId, msg.sender, price);
        return artworkId;
    }

    // 购买艺术品
    function purchaseArtwork(uint256 artworkId) public payable {
        require(_exists(artworkId), "Artwork does not exist");
        Artwork storage artwork = _artworks[artworkId];
        require(artwork.isForSale, "Artwork is not for sale");
        require(msg.value >= artwork.price, "Insufficient payment");

        // 通过调用从ERC721合约继承来的ownerOf函数，获取艺术品（根据artworkId）的当前所有者地址，存储在seller变量中
        address seller = ownerOf(artworkId);
        require(msg.sender!= seller, "Cannot buy your own artwork");

        // 转移所有权
        // 调用从ERC721合约继承来的_transfer函数，将艺术品的所有权从卖家（seller）转移到买家（msg.sender），并指定对应的艺术品ID
        _transfer(seller, msg.sender, artworkId);

        // 转移资金
        payable(seller).transfer(msg.value);

        // 更新状态
        artwork.isForSale = false;

        emit ArtworkSold(artworkId, seller, msg.sender, msg.value);
    }

    // 获取艺术品信息
    function getArtwork(uint256 artworkId) public view returns (
        string memory title,
        string memory description,
        uint256 price,
        address creator,
        uint256 createTime,
        bool isForSale
    ) {
        require(_exists(artworkId), "Artwork does not exist");
        Artwork storage artwork = _artworks[artworkId];
        return (
            artwork.title,
            artwork.description,
            artwork.price,
            artwork.creator,
            artwork.createTime,
            artwork.isForSale
        );
    }
}