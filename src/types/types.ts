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
    device_ids: Array<string>;
}

export interface Display extends BaseDocument {
    account_id: string;
    name: string;
    code: string;
    asset_contract_address: string;
    token_id: string;
    plaque_dark_mode: boolean;
    playlist_tokens: Array<AddressTokenPair>;
}

export interface Token {
    image_url: string;
    image_thumbnail_url: string;
    animation_url: string;
    name: string;
    asset_contract: AssetContract;
    token_id: string;
}

export interface AssetContract {
    address: string;
}

export interface AddressTokenPair {
    token_id: string;
    asset_contract_address: string;
}