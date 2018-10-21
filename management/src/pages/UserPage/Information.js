import React, { Component } from 'react';
import { Input } from 'antd';
import { connect } from 'dva';

import UserPage from '../../components/UserInformation/index.js';

import styles from './Information.less';

const Search = Input.Search;

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
          onSearch={value => console.log(value)}
        />
      
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
      </div>
    );
  }
}

export default connect (({ userinformation }) => ({
  information: userinformation.information,
}))(Information);
