import React from 'react';
import { useWindowDimensions } from '../../../hooks/useWindowDimensions';
import SideBarMenus from "./SideBarMenus";


const Sidebar = () => {
    const { width } = useWindowDimensions();
    if (width <= 768) {
        return null;
    }
    return <main className="sidebar"><SideBarMenus /></main>;
};

export default Sidebar;
