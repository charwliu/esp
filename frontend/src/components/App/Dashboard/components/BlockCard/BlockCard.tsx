import React, { FC } from 'react';

import s from './BlockCard.module.pcss';

interface BlockCardProps {
    overal: number | string;
    data?: number[];
    text?: string;
    color?: string;
    title: string;
}

const BlockCard: FC<BlockCardProps> = ({ overal, data, color, title, text }) => {
    return (
        <div className={s.container}>
            <div className={s.title}>{title}</div>
            <div className={s.overal}>{overal}</div>
            {data}
            {text && (
                <div>
                    {text}
                </div>
            )}
        </div>
    );
};

export default BlockCard;
