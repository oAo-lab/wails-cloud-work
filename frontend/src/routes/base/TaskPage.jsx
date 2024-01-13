import FloatChat from "../../components/FloatChat.jsx";
import {Card, Image} from "antd";

const TaskPage = () => {
    const cardData = [
        {
            title: 'Title 1',
            data: {
                info: '运行成功',
                href: 'https://picss.sunbangyan.cn/2023/12/05/4d37ff1c004f5640c47395ce3a067904.jpeg',
            }
        },
        {
            title: 'Title 2',
            data: {
                info: '',
                href: '',
            }
        },
        {
            title: 'Title 3',
            data: {
                info: '',
                href: '',
            }
        },
        {
            title: 'Title 4',
            data: {
                info: '',
                href: '',
            }
        },
    ];

    return (
        <>
            <div className='running-log'>
                <Card title="运行状态图" style={{textAlign: 'center'}}>
                    {cardData.map((item, index) => (
                        <Card.Grid id={'running-info'} key={index} title={item.title}>
                            {item.data.info}
                            {item.data.href
                                ? <Image src={item.data.href}/>
                                : <span>loading</span>
                            }
                        </Card.Grid>
                    ))}
                </Card>
            </div>
            <FloatChat/>
        </>
    );
}

export default TaskPage