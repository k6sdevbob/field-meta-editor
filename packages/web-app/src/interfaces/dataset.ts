import type { IGeoRole, ISemanticType } from "./field";


/**
 * the type given in the original data (if exists)
 */
export type IDataType = string;

export interface IDatasetFieldMeta {
    fid: string;
    name: string;
    desc: string;
    semanticType: ISemanticType;
    geoRole: IGeoRole;
    /**
     * the type given in the original data (if exists)
     */
    dataType?: IDataType;
}

export interface IDatasetInfo {
    id: string;
    name: string;
    desc: string;
    /** in seconds */
    createTime: number;
}

export interface IDataset extends IDatasetInfo {
    fieldsMeta: IDatasetFieldMeta[];
}
