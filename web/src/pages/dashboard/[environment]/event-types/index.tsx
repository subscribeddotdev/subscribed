import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutDashboard } from "@@/common/layouts/LayoutDashboard/LayoutDashboard";
import { GetServerSideProps } from "next";

interface Props {}

export default function EventTypesPage({}: Props) {
  return (
    <LayoutDashboard>
      <PageMeta title="Event types" />
      Event types
    </LayoutDashboard>
  );
}

export const getServerSideProps: GetServerSideProps = async ({ req }) => {
  return {
    props: {},
  };
};
