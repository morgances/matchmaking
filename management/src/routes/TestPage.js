import React from 'react';
import { connect } from 'dva';
import TestPageList from '../components/TestPageList';

const TestPage = ({ dispatch, testPage }) => {
  function handleDelete(id) {
    dispatch({
      type: 'testPage/delete',
      payload: id,
    });
  }
  return (
    <div>
      <h2>List of TestPage</h2>
      <TestPageList onDelete={handleDelete} testPage={testPage} />
    </div>
  );
};

export default connect(({ testPage }) => ({
  testPage,
}))(TestPage);
