import { information } from '@/services/information';

export default {
  namespace: 'userinformation',

  state: {
    information: {},
  },

  effects: {
    *userInformation({ payload }, { call, put }) {
      const response = yield call(information, payload);
      yield put({
        type: 'changeInformation',
        payload: response.data,
      });
      
      return response.data
      // Login successfully
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
