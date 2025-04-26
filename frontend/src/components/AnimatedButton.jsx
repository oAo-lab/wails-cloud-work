import React, { useState } from 'react'
import { Button } from 'antd'

const AnimatedButton = ({ buttonText, buttonType, onClick, showFlag }) => {
    const [isHovered, setIsHovered] = useState(false)

    return (
        <Button
            block
            shape="round"
            type={buttonType}
            htmlType="submit"
            style={{
                flex: 1,
                marginRight: '8px',
            }}
            onMouseEnter={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}
            onClick={onClick}
        >
            {buttonText}
            {isHovered && showFlag ? '🥕' : ''}
        </Button>
    )
}


export default AnimatedButton
