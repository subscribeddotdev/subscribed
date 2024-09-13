import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { GetServerSideProps } from "next";

interface Props {
  token: string;
}

export default function DashboardHomePage({ token }: Props) {
  return (
    <LayoutDashboard>
      <PageMeta title="Dashboard" />
      Dashboard home page
    </LayoutDashboard>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req }) => {
  return {
    props: { token: "" },
  };
};
