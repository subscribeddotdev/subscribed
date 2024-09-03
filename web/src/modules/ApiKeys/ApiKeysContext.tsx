import { ApiKey } from "@@/common/libs/backendapi/client";
import { createContext, Dispatch, PropsWithChildren, useContext, useReducer } from "react";

interface State {
  apiKeys: ApiKey[];
  environmentId: string;
}

interface Action {
  type: "set" | "remove";
  payload: ApiKey[] | string;
}

const ApiKeysContext = createContext<State>({
  apiKeys: [],
  environmentId: "",
});

const ApiKeysDispatchContext = createContext<Dispatch<Action>>((action) => {});

interface Props extends PropsWithChildren {
  initialState: State;
}

export function ApiKeysProvider({ children, initialState }: Props) {
  const [state, dispatch] = useReducer(reducer, initialState);

  return (
    <ApiKeysContext.Provider value={state}>
      <ApiKeysDispatchContext.Provider value={dispatch}>{children}</ApiKeysDispatchContext.Provider>
    </ApiKeysContext.Provider>
  );
}

export function useApiKeysContext() {
  return useContext(ApiKeysContext);
}

export function useApiKeysDispatch() {
  return useContext(ApiKeysDispatchContext);
}

function reducer(state: State, action: Action): State {
  switch (action.type) {
    case "set":
      return { ...state, apiKeys: action.payload as ApiKey[] };
    case "remove":
      return { ...state, apiKeys: state.apiKeys.filter((item) => item.id !== action.payload) };
    default:
      return state;
  }
}
