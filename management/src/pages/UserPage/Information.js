import React, { Component } from 'react';
import { Input, Button, Popover } from 'antd';
import { connect } from 'dva';

import UserPage from '../../components/UserInformation/index.js';

import styles from './Information.less';

const Search = Input.Search;

const popcontent = (
  <div>
    <p>我爱吃屎</p>
    <p>屎爱吃我</p>
  </div>
)

class Information extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div>
        <Search
          className={styles.searchFrame}
          placeholder="输入想要查询的强奸犯ID"
          enterButton
          size="large"
          onSearch={value =>console.log(value)}
        />

        <div className={styles.showDetails}>
          <div>
            <UserPage item={this.props.information[0]}/>
          </div>

          <div className={styles.buttonGroup}>
            <Popover content={popcontent} title="修改以下病句" trigger="click">
              <Button type="primary" size="large">查看联系方式</Button>
            </Popover>
          </div>
        </div>
      </div>
    );
  }
}

export default connect (({ userinformation }) => ({
  information: userinformation.information,
}))(Information);
