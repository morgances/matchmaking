import { submitInformation } from '@/services/information';

export default {
  namespace: 'information',

  state: {
    information: {},
  },

  effects: {
    *fetchInformation({ payload }, { call, put }) {
      const response = yield call(submitInformation, payload);
      console.log('后台返回的值' ,response)
      if (payload.status === 0) {
        yield put({
          type: 'changeInformation',
          payload: response.data,
        });
      }
      return response.data
    },

    reducers: {
      changeInformation(state, { payload }) {
        return {
          ...state,
          information: payload
        };
      },
    },
  }
};
