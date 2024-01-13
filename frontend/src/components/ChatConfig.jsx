import React from 'react'
import {Button, Form, Input, message, Space} from 'antd'


const ChangeConfig = ({onClose}) => {
    const [form] = Form.useForm()

    const onFinish = async (values) => {
        console.log(values)
        message.success('Submit success!')
        onClose() // 关闭模态框
    }

    const onFinishFailed = async () => {
        message.error('Submit failed!')
    }

    const onFill = () => {
        form.setFieldsValue({
            url: 'https://api.aigcbest.top/',
        })
    }

    return (
        <Form
            form={form}
            layout="vertical"
            onFinish={onFinish}
            onFinishFailed={onFinishFailed}
            autoComplete="off"
        >
            <Form.Item
                label="代理地址"
                name="url"
                rules={[
                    {required: false},
                    {type: 'url', initialValues: true},
                    {type: 'string'}
                ]}
            >
                <Input
                    allowClear
                    defaultValue="https://api.aigcbest.top/"
                />
            </Form.Item>

            <Form.Item
                label="API KEY"
                name='api_key'
                rules={[{required: true}, {type: 'string'}]}
            >
                <Input.Password
                    allowClear
                />
            </Form.Item>

            <Form.Item>
                <Space>
                    <Button type="primary" htmlType="submit">
                        Submit
                    </Button>
                    <Button htmlType="button" onClick={onFill}>
                        Fill
                    </Button>
                </Space>
            </Form.Item>
        </Form>
    )
}

export default ChangeConfig