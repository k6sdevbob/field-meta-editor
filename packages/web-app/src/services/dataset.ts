import request, { resolveServiceUrl, unwrap } from "./utils";
import type { IDataQueryWorkflowStep, IDataset, IRow } from "../interfaces";


interface IGetDatasetPayload {
    dataId: string;
}

export const getDataset = async (datasetId: string): Promise<IDataset> => {
    const url = resolveServiceUrl('/meta/query');
    const res = unwrap(await request.post<IGetDatasetPayload, IDataset>(url, { dataId: datasetId }));
    return res;
};

interface IUpdateDatasetMetaPayload {
    dataId: string;
    Meta: IDataset['meta'];
}

export const updateDatasetMeta = async (datasetId: string, fields: IDataset['meta']): Promise<IDataset['meta']> => {
    const url = resolveServiceUrl('/meta/update');
    const res = unwrap(await request.post<IUpdateDatasetMetaPayload, IDataset['meta']>(url, { dataId: datasetId, Meta: fields }));
    return res;
};

interface IGetDatasetPreviewPayload {
    dataId: string;
    payload: {
        workflow: IDataQueryWorkflowStep[];
        limit: number;
        offset: number;
    };
}

// interface IGetDatasetPreviewResult {
//     compiledSQL: string;
//     data: IRow[];
// }

export const getDatasetPreview = async (datasetId: string, limit: number): Promise<IRow[]> => {
    const url = resolveServiceUrl('/dsl/query');
    const res = unwrap(await request.post<IGetDatasetPreviewPayload, IRow[]>(url, {
        dataId: datasetId,
        payload: {
            workflow: [{
                type: 'view',
                query: [{
                    op: 'raw',
                    fields: ['*'],
                }],
            }],
            limit,
            offset: 0,
        },
    }));
    return res;
};
