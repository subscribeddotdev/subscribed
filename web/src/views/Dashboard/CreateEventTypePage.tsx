import { Alert } from "@@/common/components/Alert/Alert";
import { Button } from "@@/common/components/Button/Button";
import { Input } from "@@/common/components/Input/Input";
import { InputArea } from "@@/common/components/InputArea/InputArea";
import { PageMeta } from "@@/common/components/PageMeta/PageMeta";
import { PageTitle } from "@@/common/components/PageTitle/PageTitle";
import { apiClients, getApiError } from "@@/common/libs/backendapi/browser";
import { usePaths } from "@@/paths";
import { Flex, Grid } from "@radix-ui/themes";
import { useFormik } from "formik";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import * as yup from "yup";

export function CreateEventTypePage() {
  const paths = usePaths();
  const navigate = useNavigate();
  const [error, setError] = useState("");

  const f = useFormik({
    initialValues,
    validateOnChange: false,
    validationSchema,
    onSubmit: async (values) => {
      console.log(values);
      try {
        await apiClients().EventTypes.createEventType(values);
        navigate(paths.urls.eventTypes);
      } catch (error) {
        setError(getApiError(error));
      }
    },
  });

  return (
    <form onSubmitCapture={f.handleSubmit}>
      <PageMeta title="Create event type" />

      <Flex justify="between" mb="4">
        <PageTitle title="Create event type" />
      </Flex>

      <Flex direction="column" gap="2" mb="4">
        <Input
          required
          id="name"
          name="name"
          label="Name"
          error={f.errors.name}
          onChange={f.handleChange}
        />
        <InputArea
          id="description"
          label="Description"
          name="description"
          onChange={f.handleChange}
          error={f.errors.description}
        />
      </Flex>

      <Grid gap="2" justify="between" columns="2" width="200px" mb="4">
        <Button
          type="button"
          color="gray"
          variant="outline"
          disabled={f.isSubmitting}
          onClick={() => navigate(paths.urls.eventTypes)}
        >
          Cancel
        </Button>
        <Button
          type="submit"
          loading={f.isSubmitting}
          disabled={f.isSubmitting}
        >
          Save
        </Button>
      </Grid>

      {error && <Alert color="red">{error}</Alert>}
    </form>
  );
}

function validationSchema() {
  return yup.object().shape({
    name: yup.string().required(),
    description: yup.string().optional(),
  });
}

const initialValues = {
  name: "",
  description: "",
};
