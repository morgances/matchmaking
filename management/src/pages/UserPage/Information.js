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
          placeholder="输入想要查询的强奸犯"
          enterButton="Search"
          size="large"
          onSearch={value =>console.log(value)}
        />

        <div className={styles.showDetails}>
          <div>
            <UserPage item={this.props.information[0]}/>
            <UserPage item={this.props.information[1]}/>
            <UserPage item={this.props.information[2]}/>
            <UserPage item={this.props.information[3]}/>
            <UserPage item={this.props.information[4]}/>
            <UserPage item={this.props.information[5]}/>
            <UserPage item={this.props.information[6]}/>
            <UserPage item={this.props.information[7]}/>
            <UserPage item={this.props.information[8]}/>
            <UserPage item={this.props.information[9]}/>
            <UserPage item={this.props.information[10]}/>
            <UserPage item={this.props.information[11]}/>
            <UserPage item={this.props.information[12]}/>
            <UserPage item={this.props.information[13]}/>
            <UserPage item={this.props.information[14]}/>
            <UserPage item={this.props.information[15]}/>
            <UserPage item={this.props.information[16]}/>
            <UserPage item={this.props.information[17]}/>
          </div>

          <div className={styles.buttonGroup}>
            <Button size="large" type="primary">通过审核</Button>
            <br/>
            <Button size="large" type="primary" ghost>增加相亲次数</Button>
            <br/>
            <Button size="large" type="danger" ghost>减少相亲次数</Button>
            <br/>
            <Popover content={popcontent} title="修改以下病句" trigger="click">
              <Button size="large">查看联系方式</Button>
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
