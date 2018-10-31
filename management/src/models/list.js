import { queryList, addList, removeList } from '@/services/list';

export default {
  namespace: 'list',

  state: {
    list: [],
  },

  effects: {
    *queryList({ payload }, { call, put }) {
      const response = yield call(queryList, payload);
      console.log('打开该界面后收到的数据：', response);
      if (response.status !== 0) {
        return
      }
      yield put({
        type: 'listInformation',
        payload: response.data,
      });
    },

    *addList({ payload }, { call, put }) {
      const resp = yield call(addList, payload);
      console.log('添加之后，得到来自后台的数据：' ,resp)
      if (resp.status !== 0) {
        return false
      } else {
        const addResponse = yield call(queryList, payload);
        console.log('再次 get 的数据：', addResponse)
        if (addResponse.status !== 0) {
          return
        }
        yield put({
          type: 'addListInformation',
          payload: addResponse.data,
        });
      }
    },

    *removeList({ payload }, { call, put }) {
      const response = yield call(removeList, payload);
      if (response.status === 0) {
        const response = yield call(queryList, payload);
        if (response.status !== 0) {
          return
        }
        yield put({
          type: 'removeListInformation',
          payload: response.data,
        });
      } else {
        return false
      }
    },
  },

  reducers: {
    listInformation(state, { payload }) {
      console.log('具体数据列出：', payload)
      return {
        ...state,
        list: payload,
      };
    },

    addListInformation(state, { payload }) {
      console.log('添加后的具体数据列出：', payload)
      return {
        ...state,
        list: payload,
      };
    },

    removeListInformation(state, { payload }) {
      console.log('删除后的具体数据列出：', payload)
      return {
        ...state,
        list: payload,
      };
    },
  },
};
