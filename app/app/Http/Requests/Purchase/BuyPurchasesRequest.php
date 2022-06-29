<?php

namespace App\Http\Requests\Purchase;

use App\Http\Requests\BaseRequest;
use App\Rules\BuyingPriceRule;

class BuyPurchasesRequest extends BaseRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'buying_price'      => [
                                        'required',
                                        'numeric',
                                        'gt:0',
                                    ],
        ];
    }
}
