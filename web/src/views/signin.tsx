import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { LayoutAuth } from "@@/common/layouts/LayoutAuth/LayoutAuth";
import { SignInForm } from "@@/modules/Auth/SignInForm/SignInForm";

export default function SignInPage() {
  return (
    <LayoutAuth>
      <PageMeta title="Sign in" />
      <SignInForm />
    </LayoutAuth>
  );
}
