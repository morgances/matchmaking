import React, { Component } from 'react';
import { Input, Button, Popover } from 'antd';
import { connect } from 'dva';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import UserPage from '../../components/UserInformation/index.js';

import styles from './Information.less';

const Search = Input.Search;

class Information extends Component {
  constructor(props) {
    super(props);
    this.state = {
      value: '',
      information: {},
    };
  }

  handleChange = (input) => {
    this.setState({
      value: input
    })
    const data = {
      "target_open_id": input
    }
    const { dispatch } = this.props;
    dispatch({
      type: 'userinformation/userInformation',
      payload: {
        ...data,
      },
    }).then((result) => {
      this.setState({
        information: result
      })
    }).catch((err) => {
      console.log(err);
      return false
    });
  }
  
  render() {
    return (
      <PageHeaderWrapper>
        <Search
          ref="search"
          className={styles.searchFrame}
          placeholder="输入想要查询的用户ID"
          enterButton
          size="large"
          onSearch={this.handleChange}
        />

        <div className={styles.selfComponent}>
          <UserPage item={this.state.information} />
        </div>
      </PageHeaderWrapper>
    );
  }
}

export default connect (({ userinformation }) => ({
  information: userinformation.information,
}))(Information);
