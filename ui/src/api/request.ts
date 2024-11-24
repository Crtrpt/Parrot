import queryString from "query-string";

export class Resp {
    Code!: Number;
    Msg!: String;
    Data:any;
}

let endpoint = import.meta.env.VITE_API_ENDPOINT;
export const GET = async (path: string, query: any):Promise<Resp> => {
  return request({
    method: "get",
    path: path,
    query: query,
  });
};

export const POST = async (path: string, body: any, query: any):Promise<Resp> => {
  return request({ method: "post", path: path, query: query, body: body });
};

const request = async (req: {
  method: string;
  path: string;
  query?: any;
  header?: any | undefined;
  body?: any;
  opt?: any;
}):Promise<Resp> => {
  return fetch(endpoint + req.path + "?" + queryString.stringify(req.query), {
        ...req.opt,
        ...{ method: req.method },
        ...{
            header: req.header,
        },
        ...{
            body: JSON.stringify(req.body),
        }
    }).then((response) => response.json());
};

