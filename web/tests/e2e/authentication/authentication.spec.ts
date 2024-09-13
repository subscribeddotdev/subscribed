import { faker } from "@faker-js/faker";
import { expect, test } from "@playwright/test";

test("Authentication/SignUp", async ({ page }) => {
  const user = {
    firstName: faker.person.firstName(),
    lastName: faker.person.lastName(),
    email: "",
    password: faker.internet.password(),
  };

  user.email = faker.internet
    // Include the first and last name in the email address
    .email({ firstName: user.firstName, lastName: user.lastName })
    .toLowerCase()
    // Add timestamp to avoid collision
    .replace("@", `+${new Date().getTime()}@`);

  await page.goto("http://localhost:3000/signup");
  await page.getByTestId("SignUpForm_Inp_FirstName").fill(user.firstName);
  await page.getByTestId("SignUpForm_Inp_LastName").fill(user.lastName);
  await page.getByTestId("SignUpForm_Inp_Email").fill(user.email);
  await page.getByTestId("SignUpForm_Inp_Password").fill(user.password);
  await page.getByTestId("SignUpForm_Btn_CreateAccount").click();
  await page.waitForURL(/signin/);
  await expect(
    page.getByTestId("SignInForm_Alert_AccountCreated"),
  ).toBeVisible();
});
