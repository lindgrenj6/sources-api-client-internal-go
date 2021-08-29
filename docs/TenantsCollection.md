# TenantsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**CollectionMetadata**](CollectionMetadata.md) |  | [optional] 
**Links** | Pointer to [**CollectionLinks**](CollectionLinks.md) |  | [optional] 
**Data** | Pointer to [**[]Tenant**](Tenant.md) |  | [optional] 

## Methods

### NewTenantsCollection

`func NewTenantsCollection() *TenantsCollection`

NewTenantsCollection instantiates a new TenantsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantsCollectionWithDefaults

`func NewTenantsCollectionWithDefaults() *TenantsCollection`

NewTenantsCollectionWithDefaults instantiates a new TenantsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *TenantsCollection) GetMeta() CollectionMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TenantsCollection) GetMetaOk() (*CollectionMetadata, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TenantsCollection) SetMeta(v CollectionMetadata)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TenantsCollection) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetLinks

`func (o *TenantsCollection) GetLinks() CollectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TenantsCollection) GetLinksOk() (*CollectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TenantsCollection) SetLinks(v CollectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TenantsCollection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetData

`func (o *TenantsCollection) GetData() []Tenant`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TenantsCollection) GetDataOk() (*[]Tenant, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TenantsCollection) SetData(v []Tenant)`

SetData sets Data field to given value.

### HasData

`func (o *TenantsCollection) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


