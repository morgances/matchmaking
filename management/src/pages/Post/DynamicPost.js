import React, { Component } from 'react';
import { connect } from 'dva';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';

import Post from '../../components/Post/index';

class DynamicPost extends Component {
  constructor(props) {
    super(props);
  }

  getData = async() => {
    const { dispatch } = this.props;
    await dispatch({
      type: 'post/fetchPost',
    })
  }

  componentDidMount() {
    this.getData();
  }

  render() {
    const dispatch = this.props.dispatch;

    return (
      <PageHeaderWrapper>
        {
          this.props.post.map((item) => {
            item.dispatch = dispatch;
            return (
              <Post
                key={item.id}
                item={item}
              />
            );
          })
        }
      </PageHeaderWrapper>
    );
  }
}

export default connect(({ post }) => ({
  ...post,
}))(DynamicPost);
