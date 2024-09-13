import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";

export default function DashboardHomePage() {
  return (
    <LayoutDashboard>
      <PageMeta title="Dashboard" />
      Dashboard home page
    </LayoutDashboard>
  );
}
