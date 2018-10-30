import React, { Component } from 'react';
import { connect } from 'dva';
import { Card, Table, Divider, Button } from 'antd';

import PageHeaderWrapper from '@/components/PageHeaderWrapper';

import styles from './BasicProfile.less';

@connect(({ profile }) => ({
  ...profile,
}))

class BasicProfile extends Component {
  constructor(props) {
    super(props);
  }

  success = (id) => {
    const { dispatch } = this.props;
    dispatch({
      type: 'profile/successProfile',
      payload: {
        target_id: id,
      }
    })
  }

  failure = (id) => {
    const { dispatch } = this.props;
    dispatch({
      type: 'profile/failureProfile',
      payload: {
        target_id: id,
      }
    })
  }

  componentDidMount() {
    this.getData();
  }

  getData = async () => {
    const { dispatch } = this.props;
    await dispatch({
      type: 'profile/fetchProfile',
    })
  }

  render() {
    const columns = [{
        title: '订单ID',
        key: 'id',
        dataIndex: 'id',
      },{
        title: '用户ID',
        key: 'open_id',
        dataIndex: 'open_id',
      },{
        title: '商品ID',
        key: 'goods_id',
        dataIndex: 'goods_id',
      },{
        title: '买家姓名',
        key: 'buyer_name',
        dataIndex: 'buyer_name',
      },{
        title: '商品名称',
        key: 'goods_name',
        dataIndex: 'goods_name',
      },{
        title: '兑换时间',
        key: 'date_time',
        dataIndex: 'date_time',
      },{
        title: '使用积分',
        key: 'cost',
        dataIndex: 'cost',
      },{
        title: '订单状态',
        key: 'finished',
        dataIndex: 'finished',
        render: (text, record) => (
          <div>
            <Button
              onClick={() => this.success(record.id)}
              type="primary"
            >
              确认
            </Button>

            <Divider type="vertical" />

            <Button
              onClick={() => this.failure(record.id)}
              type="danger"
            >
              取消
            </Button>
          </div>
        )
      }
    ];

    return (
      <PageHeaderWrapper>
        <Card bordered>
          <Divider style={{ marginBottom: 32 }} />

          <div className={styles.title}>积分交易处理</div>

          <Table
            style={{ marginBottom: 16 }}
            columns={columns}
            dataSource={this.props.basicGoods}
          />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default BasicProfile;
