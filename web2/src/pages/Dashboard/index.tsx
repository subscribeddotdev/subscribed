import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { Outlet } from "react-router-dom";

export default function DashboardRoot() {
  return (
    <LayoutDashboard>
      <Outlet />
    </LayoutDashboard>
  );
}
