const basicGoods = [
  {
    id: '1234561',
    name: '矿泉水 550ml',
    barcode: '12421432143214321',
    price: '2.00',
    num: '1',
    amount: '2.00',
  },
  {
    id: '1234562',
    name: '凉茶 300ml',
    barcode: '12421432143214322',
    price: '3.00',
    num: '2',
    amount: '6.00',
  },
  {
    id: '1234563',
    name: '好吃的薯片',
    barcode: '12421432143214323',
    price: '7.00',
    num: '4',
    amount: '28.00',
  },
  {
    id: '1234564',
    name: '特别好吃的蛋卷',
    barcode: '12421432143214324',
    price: '8.50',
    num: '3',
    amount: '25.50',
  },
];

const basicProgress = [
  {
    key: '1',
    id: '00000001',
    open_id: 'wechat-00000001',
    goods_id: '千万分之一',
    buyer_name: '任我行',
    goods_name: '矿泉水',
    date_time: '2017-10-01 14:10',
    cost: '10',
    finished: 'processing',
  },
  {
    key: '2',
    id: '000000002',
    open_id: 'wechat-00000002',
    goods_id: '千万分之二',
    buyer_name: '我是王八羔子',
    goods_name: '可乐',
    date_time: '2018-01-01 18:10',
    cost: '30',
    finished: 'processing',
  },
  {
    key: '3',
    id: '000000003',
    open_id: 'wechat-00000003',
    goods_id: '千万分之三',
    buyer_name: '我是鸡掰',
    goods_name: '雪花',
    date_time: '2018-08-30 10:19',
    cost: '40',
    finished: 'processing',
  },
  {
    key: '4',
    id: '000000004',
    open_id: 'wechat-00000004',
    goods_id: '千万分之四',
    buyer_name: '独孤求败',
    goods_name: '红糖',
    date_time: '2018-03-07 09:20',
    cost: '100',
    finished: 'success',
  },
  {
    key: '5',
    id: '000000005',
    open_id: 'wechat-00000005',
    goods_id: '千万分之五',
    buyer_name: '白起',
    goods_name: '花生米',
    date_time: '2018-09-09 09:09',
    cost: '1',
    finished: 'success',
  },
  {
    key: '6',
    id: '000000006',
    open_id: 'wechat-00000006',
    goods_id: '千万分之六',
    buyer_name: '令狐冲',
    goods_name: '桃子',
    date_time: '2018-10-10 10:10',
    cost: '5',
    finished: 'success',
  },
];

const advancedOperation1 = [
  {
    key: 'op1',
    type: '订购关系生效',
    name: '曲丽丽',
    status: 'agree',
    updatedAt: '2017-10-03  19:23:12',
    memo: '-',
  },
  {
    key: 'op2',
    type: '财务复审',
    name: '付小小',
    status: 'reject',
    updatedAt: '2017-10-03  19:23:12',
    memo: '不通过原因',
  },
  {
    key: 'op3',
    type: '部门初审',
    name: '周毛毛',
    status: 'agree',
    updatedAt: '2017-10-03  19:23:12',
    memo: '-',
  },
  {
    key: 'op4',
    type: '提交订单',
    name: '林东东',
    status: 'agree',
    updatedAt: '2017-10-03  19:23:12',
    memo: '很棒',
  },
  {
    key: 'op5',
    type: '创建订单',
    name: '汗牙牙',
    status: 'agree',
    updatedAt: '2017-10-03  19:23:12',
    memo: '-',
  },
];

const advancedOperation2 = [
  {
    key: 'op1',
    type: '订购关系生效',
    name: '曲丽丽',
    status: 'agree',
    updatedAt: '2017-10-03  19:23:12',
    memo: '-',
  },
];

const advancedOperation3 = [
  {
    key: 'op1',
    type: '创建订单',
    name: '汗牙牙',
    status: 'agree',
    updatedAt: '2017-10-03  19:23:12',
    memo: '-',
  },
];

const getProfileBasicData = {
  basicGoods,
  basicProgress,
};

const getProfileAdvancedData = {
  advancedOperation1,
  advancedOperation2,
  advancedOperation3,
};

export default {
  'GET /api/profile/advanced': getProfileAdvancedData,
  'GET /api/profile/basic': getProfileBasicData,
};
