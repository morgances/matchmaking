import React, { Component } from 'react';
import { connect } from 'dva';
import { Card, Table, Button, Divider, Input, Form, Modal, Popconfirm } from 'antd';

import PageHeaderWrapper from '@/components/PageHeaderWrapper';

// const Search = Input.Search;

const FormItem = Form.Item;

const CollectionCreateForm = Form.create()(
  class extends React.Component {
    render() {
      const { visible, onCancel, onCreate, form } = this.props;

      const { getFieldDecorator } = form;

      return (
        <Modal
          visible={visible}
          onCancel={onCancel}
          onOk={onCreate}
          title="创建商品信息"
          okText="确认"
        >
          <Form layout="vertical">
            <FormItem label="商品标题">
              {getFieldDecorator('title')(<Input placeholder="请输入标题" />)}
            </FormItem>

            <FormItem label="商品价格">
              {getFieldDecorator('price')(<Input placeholder="请输入价格"/>)}
            </FormItem>

            <FormItem label="商品描述">
              {getFieldDecorator('description')(<Input type="textarea" placeholder="请输入描述内容" />)}
            </FormItem>
          </Form>
        </Modal>
      );
    }
  }
);

@connect(({ list }) => ({
  ...list,
}))

class BasicList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      visible: false,
    };
  }

  componentDidMount() {
    this.getData();
  }

  getData = async () => {
    const { dispatch } = this.props;
    await dispatch({
      type: 'list/queryList',
    })
  }

  showModal = () => {
    this.setState({
      visible: true,
    })
  }

  handleCancel = () => {
    this.setState({
      visible: false,
    });
  }

  handleCreate = () => {
    const form = this.formRef.props.form;
    form.validateFields(async (err, values) => {
      if (err) {
        return;
      }
      console.log('给后台传入的数据：', values);
      values.price = +values.price
      const { dispatch } = this.props;
      await dispatch({
        type: 'list/addList',
        payload: {
          ...values,
        }
      })
      form.resetFields();
      this.setState({
        visible: false,
      });
    });
  }

  saveFormRef = (formRef) => {
    this.formRef = formRef;
  }

  deleteConfirm = (itemId) => {
    const { dispatch } = this.props;
    dispatch({
      type: 'list/removeList',
      payload: {
        target_id: itemId,
      },
    })
  }

  render() {
    const columns = [{
      title: '商品 ID',
      dataIndex: 'id',
    },{
      title: '商品标题',
      dataIndex: 'title',
    },{
      title: '商品价格',
      dataIndex: 'price',
    },{
      title: '商品描述',
      dataIndex: 'description',
    },{
      title: '操作',
      dataIndex: 'operation',
      render: (text, record) => (
        this.props.list.length >= 1
        ? (
          <Popconfirm
            title="确定删除？"
            onConfirm={() => this.deleteConfirm(record.id)}
            okText="确认"
            cancelText="取消"
          >
            <Button type="danger" ghost>删除</Button>
          </Popconfirm>
        ) : null
      )
    }];

    return (
      <PageHeaderWrapper>
        <Card bordered>
          {/* <Search
            enterButton
            placeholder="输入商品名称"
          /> */}
          <h1>标准商品列表</h1>

          <Divider style={{ marginBottom: 32 }} />

          <div>
            <Button
              block
              icon="plus"
              type="primary"
              size="large"
              onClick={this.showModal}
            >
              添加商品
            </Button>

            <CollectionCreateForm
              wrappedComponentRef={this.saveFormRef}
              visible={this.state.visible}
              onCancel={this.handleCancel}
              onCreate={this.handleCreate}
            />
          </div>

          <br/>

          <Table
            style={{ marginBottom: 16 }}
            columns={columns}
            dataSource={this.props.list}
          />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default BasicList;
