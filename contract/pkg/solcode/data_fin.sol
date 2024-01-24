// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

/**
 * @title Owner
 * @dev Set & change owner
 */
contract Owner {

    address private owner;

    // event for EVM logging
    event OwnerSet(address indexed oldOwner, address indexed newOwner);

    // modifier to check if caller is owner
    modifier isOwner() {
        // If the first argument of 'require' evaluates to 'false', execution terminates and all
        // changes to the state and to Ether balances are reverted.
        // This used to consume all gas in old EVM versions, but not anymore.
        // It is often a good idea to use 'require' to check if functions are called correctly.
        // As a second argument, you can also provide an explanation about what went wrong.
        require(msg.sender == owner, "Caller is not owner");
        _;
    }

    /**
     * @dev Set contract deployer as owner
     */
    constructor() {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        emit OwnerSet(address(0), owner);
    }

    /**
     * @dev Change owner
     * @param newOwner address of new owner
     */
    function ChangeOwner(address newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    /**
     * @dev Return owner address 
     * @return address of owner
     */
    function GetOwner() public view returns (address) {
        return owner;
    }
}

// 管理员功能接口
interface IAdmin{
    function AddAdmin(address admin,string memory info) external;
    function DeleteAdmin(address admin) external;
    function SetEnableAdmin(address admin,bool enable) external;
    function IsAdminEnable(address admin) external view returns(bool);
}

// 管理员合约
contract Admin is IAdmin,Owner{
    struct AdminInfo{
        bool enable;
        string info;
        bool isValid;
    }
    mapping(address => AdminInfo) adminMap;
    address[] adminArr;
    uint32 adminTotal;
    
    // 初始化合约,将初始化地址设为管理员之一
    constructor() Owner() {
        adminMap[msg.sender] = AdminInfo(true,"root admin",true);
        adminArr.push(msg.sender);
        adminTotal++;
    }

    // 添加
    function AddAdmin(address admin,string memory info) override  external isOwner{
        require(!adminMap[admin].enable,"cannot change info of admin in using");
        adminMap[admin] = AdminInfo(true,info,true);
        adminArr.push(admin);
        adminTotal++;
    }

    // 设置可用性
    function SetEnableAdmin(address admin,bool enable) override external isOwner{
        adminMap[admin].enable=enable;
    }

    // 查看可用性
    function IsAdminEnable(address admin) override external view returns(bool){
        return adminMap[admin].enable;
    }
    
    // 获取管理员接口
    function GetAdminInfos() external view returns(address[] memory,AdminInfo[] memory){
        AdminInfo[] memory rets=new AdminInfo[](adminTotal);
        address[] memory addrs=new address[](adminTotal);
        
        uint j=0;
        for (uint i=0;i<adminArr.length;i++){
            if(adminArr[i]==address(0)){
                continue;
            }
            if ( adminMap[adminArr[i]].isValid){
                rets[j]=adminMap[adminArr[i]];
                addrs[j]=adminArr[i];
                j++;
            }
        }
        return (addrs,rets);
    }

    // 删除管理员
    function DeleteAdmin(address admin) override external isOwner{
        // require(!adminMap[admin].enable,"cannot delete admin in using");
        delete adminMap[admin];
        adminTotal--;
        for (uint i=0;i<adminArr.length;i++){
            if (admin==adminArr[i]){
                adminArr[i]=address(0);
            }
        }
    }

    modifier isAdmin() {
        require(adminMap[msg.sender].enable || msg.sender == GetOwner(),"Caller is not admin");
        _;
    }
}

contract DataFin is Admin{
    struct FinInfo{
        string ID;
        string Name;
        string Info;
        bool isValid;
    }

    struct IDFinInfo{
        string ID;
        string Name;
        string Info;
        bool ChangeAble;
        bool isValid;
    }

    // finInfo.id => finInfo
    mapping(string => FinInfo) finInfoMap;
    string[] finInfoArr;
    // finIDInfo.id => IDFinInfo
    mapping(string => IDFinInfo) idFinInfoMap;
    string[] idFinInfoArr;

    // finInfo.id => (sha256=>uinx)
    // mapping(string => mapping(uint256=>uint64)) finKeys;
    mapping(string => mapping(uint256=>uint256)) finKeys;
   
    // finInfo.id => ((info.id)=>sha224+uinx32)
    // mapping(string => mapping(uint128=>uint256)) idFinKeys;
    mapping(string => mapping(uint64=>uint256)) idFinKeys;
    mapping(string => mapping(uint64=>uint256)) idFinKeysTime;

    string version="DataFin-v1.2.0-with-Admin-Owner";
    
    // event for EVM logging
    event DataFinCreate(address indexed cteator);

    // event for EVM logging
    event TopicCreate(address indexed cteator);

    constructor() Admin() {
        emit DataFinCreate(msg.sender);
    }

    function CreateTopic(string memory topicID,string memory name,string memory info) external isAdmin() {
        require(!finInfoMap[topicID].isValid,"topic is already exist");
        finInfoMap[topicID] = FinInfo(topicID,name,info,true);
        finInfoArr.push(topicID);
    }
    
    function CreateIDTopic(string memory topicID,string memory name,string memory info,bool changeAble) external isAdmin() {
        require(!idFinInfoMap[topicID].isValid,"topic is already exist");
        idFinInfoMap[topicID] = IDFinInfo(topicID,name,info,changeAble,true);
        idFinInfoArr.push(topicID);
    }

    function AddItems(string memory topicID,uint256[] memory vals)external isAdmin() returns (bool[] memory){
        require(finInfoMap[topicID].isValid,"topic is not exist");

        bool [] memory rets=new bool[](vals.length);
        for(uint i = 0; i < vals.length; i++) {
            if (finKeys[topicID][vals[i]]>0){
                rets[i]=false;
            }else{
                rets[i]=true;
                finKeys[topicID][vals[i]]=block.timestamp;
            }
        }
        return rets;
    }

    function AddIDsItems(string memory topicID,uint64[] memory ids,uint256[] memory vals)external isAdmin() returns (bool[] memory){
        require(idFinInfoMap[topicID].isValid,"topic is not exist");
        require(ids.length==vals.length,"the lengths of the two are different");
        bool [] memory rets=new bool[](ids.length);
        for(uint i = 0; i < ids.length; i++) {
            if (idFinInfoMap[topicID].ChangeAble){
                rets[i]=true;
                idFinKeys[topicID][ids[i]]=vals[i];
                idFinKeysTime[topicID][ids[i]]=block.timestamp;
                continue ;
            }
            if (idFinKeys[topicID][ids[i]] > 0){
                rets[i]=false;
            }else{
                rets[i]=true;
                idFinKeys[topicID][ids[i]]=vals[i];
                idFinKeysTime[topicID][ids[i]]=block.timestamp;
            }
        }
        return rets;
    }

    // if the rets[i] > 0 ,it`s true
    function VerifyItems(string memory topicID,uint256[] memory vals) view external isAdmin() returns (uint[] memory){
        require(finInfoMap[topicID].isValid,"topic is not exist");

        uint [] memory rets=new uint[](vals.length);
        for(uint i = 0; i < vals.length; i++) {
            rets[i]=finKeys[topicID][vals[i]];
        }
        return rets;
    }

    function VerifyIDItems(string memory topicID,uint64[] memory ids,uint256[] memory vals) view external isAdmin() returns (uint[] memory){
        require(idFinInfoMap[topicID].isValid,"topic is not exist");
        require(ids.length==vals.length,"the lengths of the two are different");
        uint [] memory rets=new uint[](ids.length);
        for(uint i = 0; i < ids.length; i++) {
            if (idFinKeys[topicID][ids[i]] == vals[i]){
                rets[i]=idFinKeysTime[topicID][ids[i]];
            }
        }
        return rets;
    }

    function Version(string memory _version) view external returns (bool ret){
        bytes memory v1 = bytes(_version);
        bytes memory v2 = bytes(version);
        // 如果长度不等，直接返回
        if (v1.length != v2.length) return false;
        // 按位比较
        for(uint i = 0; i < v1.length; i ++) {
            if(v1[i] != v2[i]) return false;
        }
    }

    function GetTopics()view external returns (FinInfo[] memory) {
        FinInfo[] memory rets=new FinInfo[](finInfoArr.length);
        for(uint i=0 ; i< finInfoArr.length; i++){
            rets[i]=finInfoMap[finInfoArr[i]];
        }
        return rets;
    }

    function GetIDTopics()view external returns (IDFinInfo[] memory) {
        IDFinInfo[] memory rets=new IDFinInfo[](idFinInfoArr.length);
        for(uint i=0 ; i< idFinInfoArr.length; i++){
            rets[i]=idFinInfoMap[idFinInfoArr[i]];
        }
        return rets;
    }
}
