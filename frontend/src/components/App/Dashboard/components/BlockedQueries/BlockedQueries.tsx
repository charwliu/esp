import theme from 'Lib/theme';
import React, { FC, useContext, useState } from 'react';

import Store from 'Store';

import s from './BlockedQueries.module.pcss';

interface BlockCardProps {
    ads: number;
    trackers: number;
    other: number;
}

const renderActiveShape = (props: any): any => {
    const {
        cx, cy, innerRadius, outerRadius, startAngle, endAngle,
        fill, payload, percent,
    } = props;
    return (
        <g>
            <text x={cx} y={cy - 11} dy={8} textAnchor="middle" fill={fill}>{payload.name}</text>
            <text x={cx} y={cy + 18} dy={8} fontSize={24} textAnchor="middle" >{Math.round(percent * 100)}%</text>
        </g>
    );
};

const BlockedQueries: FC<BlockCardProps> = ({ ads, trackers, other }) => {
    const store = useContext(Store);
    const [activeIndex, setActiveIndex] = useState(0);
    const { ui: { intl } } = store;
    const data = [
        { name: intl.getMessage('other'), value: other, color: theme.chartColors.gray700 },
        { name: intl.getMessage('ads'), value: ads, color: theme.chartColors.red },
        { name: intl.getMessage('trackers'), value: trackers, color: theme.chartColors.orange },
    ];
    const onChart: any = (_: any, index: number) => {
        setActiveIndex(index);
    };
    return (
        <div className={s.container}>
            <div className={s.title}>{intl.getMessage('dashboard_blocked_queries')}</div>
            <div className={s.pie}>
            </div>
        </div>
    );
};

export default BlockedQueries;
