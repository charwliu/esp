import cn from 'classnames';
import React, {FC} from 'react';

import {IconType} from '@components/Icons';

import s from './Icon.module.less';

interface IconProps {
    icon: IconType;
    color?: string;
    className?: string;
    onClick?: () => void;
}

const Icon: FC<IconProps> = ({icon, color, className, onClick}) => {
    const iconClass = cn(s.icon, color, className);
    return (
        <svg className={iconClass} onClick={onClick}>
            <use xlinkHref={`#${icon}`} />
        </svg>
    );
};

export default Icon;
export type {IconType} from '@components/Icons';
