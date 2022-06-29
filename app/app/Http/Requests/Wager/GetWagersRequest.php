<?php

namespace App\Http\Requests\Wager;

use App\Http\Requests\BaseRequest;

class GetWagersRequest extends BaseRequest
{
    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'limit' => 'numeric',
            'page'  => 'numeric',
        ];
    }
}
