import firebase from "firebase";

export interface BaseDocument {
    created_at: firebase.firestore.Timestamp;
    updated_at: firebase.firestore.Timestamp;
}

export class BaseEntity {

    static createBaseEntity(): BaseDocument {
        const now = firebase.firestore.Timestamp.now();

        return { created_at: now, updated_at: now };
    }

    static updateBaseEntity(): Partial<BaseDocument> {
        const now = firebase.firestore.Timestamp.now();

        return { updated_at: now }
    }
}

export interface FirestoreDocument<Entity extends BaseDocument> {
    entity: Entity;
    id: string;
}

export interface Account extends BaseDocument {
    address: string;
    signature: string; 
    admin: boolean;
    saved_display_ids: Array<string>; // if user is an admin, devices can be saved here and remembered even if another user has assumed control of the device
}

export interface Display extends BaseDocument {
    account_id: string;
    name: string;
    code: string;
    asset_contract_address: string;
    token_id: string;
    plaque_dark_mode: boolean;
    playlist_tokens: Array<TokenMeta>;
}

export interface OpenseaToken {
    image_url: string;
    image_thumbnail_url: string;
    animation_url: string;
    name: string;
    asset_contract: AssetContract;
    token_id: string;
    description: string;
    permalink: string;
    creator: any;
    orders: Array<any>;
}

export interface AssetContract {
    address: string;
}

export interface TokenMeta extends BaseDocument {
    token_id: string;
    asset_contract_address: string;
    tag: string;
    name: string;
    // created_at: string;
}