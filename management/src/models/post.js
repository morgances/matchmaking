import { queryFakeList, removeFakeList, addFakeList, updateFakeList } from '@/services/api';

export default {
  namespace: 'post',

  state: {
    dynamic: [{
      open_id: '001',
      avatar: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
      title: '喜欢我吗',
      content: '技术猫是最棒的交友空间了，我超喜欢的',
      img: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    }, {
      open_id: '002',
      avatar: 'http://m1.aboluowang.com/uploadfile/2017/0901/20170901042508280.jpg',
      title: '小哥哥们，晚上好',
      content: '4000年美女鞠婧祎证件照曝光青春靓丽',
      img: 'http://m1.aboluowang.com/uploadfile/2017/0901/20170901042508280.jpg',
    }, {
      open_id: '003',
      avatar: 'https://img.gq.com.tw/_rs/645/userfiles/sm/sm1024_images_A1/34940/2018011671096217.jpg',
      title: '软妹报道',
      content: '最讓觀眾分神的10位新生代美女主播',
      img: 'https://img.gq.com.tw/_rs/645/userfiles/sm/sm1024_images_A1/34940/2018011671096217.jpg',
    }, {
      open_id: '004',
      avatar: 'http://m1.ablwang.com/uploadfile/2017/0901/20170901042510839.jpg',
      title: '你最爱的人是我',
      content: '韓國最美的美女她無可挑剔的容顏能將人融化',
      img: 'http://m1.ablwang.com/uploadfile/2017/0901/20170901042510839.jpg',
    }, {
      open_id: '005',
      avatar: 'http://www.5didao.com/wp-content/uploads/2017/07/Flamingo-Beach-2-2-1280x640.jpg',
      title: '喜欢我的可以加个关注哦',
      content: '我是大四的基佬，喜欢搞基，有意者私聊，猛戳→→微信号：技术猫',
      img: 'http://www.5didao.com/wp-content/uploads/2017/07/Flamingo-Beach-2-2-1280x640.jpg',
    }, {
      open_id: '006',
      avatar: 'http://5b0988e595225.cdn.sohucs.com/images/20171028/26d0a9838fcc444e8e5ce7c5346ec5a2.jpeg',
      title: '么么哒',
      content: '气质大放送',
      img: 'http://5b0988e595225.cdn.sohucs.com/images/20171028/26d0a9838fcc444e8e5ce7c5346ec5a2.jpeg',
    }, {
      open_id: '007',
      avatar: 'http://www.fullyu.com/ueditor/php/upload/image/20161003/1475480755266911.jpg',
      title: '你来，我不让你走',
      content: '美女与野兽',
      img: 'http://www.fullyu.com/ueditor/php/upload/image/20161003/1475480755266911.jpg',
    }],
  },

  effects: {
    *fetch({ payload }, { call, put }) {
      const response = yield call(queryFakeList, payload);
      yield put({
        type: 'queryList',
        payload: Array.isArray(response) ? response : [],
      });
    },

    *appendFetch({ payload }, { call, put }) {
      const response = yield call(queryFakeList, payload);
      yield put({
        type: 'appendList',
        payload: Array.isArray(response) ? response : [],
      });
    },

    *submit({ payload }, { call, put }) {
      let callback;
      if (payload.id) {
        callback = Object.keys(payload).length === 1 ? removeFakeList : updateFakeList;
      } else {
        callback = addFakeList;
      }
      const response = yield call(callback, payload); // post
      yield put({
        type: 'queryList',
        payload: response,
      });
    },
  },

  reducers: {
    queryList(state, action) {
      return {
        ...state,
        list: action.payload,
      };
    },

    appendList(state, action) {
      return {
        ...state,
        list: state.list.concat(action.payload),
      };
    },
  },
};
