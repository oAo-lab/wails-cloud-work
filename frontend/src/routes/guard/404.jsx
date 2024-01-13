import React, {useEffect} from 'react';
import {Link} from 'react-router-dom';
import "../../css/NotFoundPage.css"

const NotFoundPage = () => {
    // 添加一些JS动画效果
    useEffect(() => {
        const animateText = () => {
            const text = document.querySelector('.not-found-text');
            text.classList.add('animated', 'bounce')
        };
        animateText();
    }, []);

    return (
        <div className='not-body'>
            <div className="not-found-container">
                <h1 className="not-found-text">404 - Page Not Found</h1>
                <p className="not-found-description">
                    Oops! The page you are looking for doesn't exist.
                </p>
                <Link to="/" className="back-to-home">
                    Back to Home
                </Link>
            </div>
        </div>
    );
};

export default NotFoundPage;
