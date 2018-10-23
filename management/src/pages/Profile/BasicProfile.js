import React, { Component } from 'react';
import { connect } from 'dva';
import { Card, Badge, Table, Divider } from 'antd';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import styles from './BasicProfile.less';

const progressColumns = [
  {
    title: '订单ID',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '用户ID',
    dataIndex: 'open_id',
    key: 'open_id',
  },
  {
    title: '商品ID',
    dataIndex: 'goods_id',
    key: 'goods_id',
  },
  {
    title: '买家姓名',
    dataIndex: 'buyer_name',
    key: 'buyer_name',
  },
  {
    title: '商品名称',
    dataIndex: 'goods_name',
    key: 'goods_name',
  },
  {
    title: '兑换时间',
    dataIndex: 'date_time',
    key: 'date_time',
  },
  {
    title: '使用积分',
    dataIndex: 'cost',
    key: 'cost',
  },
  {
    title: '订单状态',
    dataIndex: 'finished',
    key: 'finished',
    render: text =>
      text === 'success' ? (
        <Badge status="success" text="成功" />
      ) : (
        <Badge status="processing" text="进行中" />
      ),
  },
];

@connect(({ profile, loading }) => ({
  profile,
  loading: loading.effects['profile/fetchBasic'],
}))

class BasicProfile extends Component {
  componentDidMount() {
    const { dispatch } = this.props;
    dispatch({
      type: 'profile/fetchBasic',
    });
  }

  render() {
    const { profile, loading } = this.props;
    const { basicGoods, basicProgress } = profile;
    let goodsData = [];
    if (basicGoods.length) {
      let num = 0;
      let amount = 0;
      basicGoods.forEach(item => {
        num += Number(item.num);
        amount += Number(item.amount);
      });
      goodsData = basicGoods.concat({
        id: '总计',
        num,
        amount,
      });
    }

    const renderContent = (value, row, index) => {
      const obj = {
        children: value,
        props: {},
      };
      if (index === basicGoods.length) {
        obj.props.colSpan = 0;
      }
      return obj;
    };

    return (
      <PageHeaderWrapper>
        <Card bordered={false}>
          <Divider style={{ marginBottom: 32 }} />
          <div className={styles.title}>积分交易处理</div>
          <Table
            style={{ marginBottom: 16 }}
            pagination={false}
            loading={loading}
            dataSource={basicProgress}
            columns={progressColumns}
          />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default BasicProfile;
