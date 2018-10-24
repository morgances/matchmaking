import React, { Component } from 'react';
import { Tag, Button, Avatar, Modal } from 'antd';
import { connect } from 'dva';

import Post from '../../components/Post/index';
import styles from './DynamicPost.less';

const confirm = Modal.confirm;

class DynamicPost extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  showConfirm() {
    confirm({
      title: 'Are you sure delete this task?',
      content: 'Some descriptions',
      okText: 'Yes',
      okType: 'danger',
      cancelText: 'No',
      onOk() {
        console.log('OK');
      },
      onCancel() {
        console.log('Cancel');
      },
    });
  }

  render() {
    return (
      <div>
        <Post item={this.props.dynamic[0]} />
        <Post item={this.props.dynamic[1]} />
        <Post item={this.props.dynamic[2]} />
        <Post item={this.props.dynamic[3]} />
        <Post item={this.props.dynamic[4]} />
        <Post item={this.props.dynamic[5]} />
        <Post item={this.props.dynamic[6]} />
      </div>
    );
  }
}

export default connect(({ post }) => ({
  dynamic: post.dynamic,
}))(DynamicPost);
