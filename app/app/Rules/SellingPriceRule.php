<?php

namespace App\Rules;

use Illuminate\Contracts\Validation\Rule;

class SellingPriceRule implements Rule
{

    protected $totalWagerValue;
    protected $sellingPercentage;

    /**
     * Create a new rule instance.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }

    public function totalWagerValue($totalWagerValue)
    {
        $this->totalWagerValue = $totalWagerValue;

        return $this;
    }

    public function sellingPercentage($sellingPercentage)
    {
        $this->sellingPercentage = $sellingPercentage;

        return $this;
    }

    /**
     * Determine if the validation rule passes.
     *
     * @param  string  $attribute
     * @param  mixed  $value
     * @return bool
     */
    public function passes($attribute, $value)
    {
        $gtSellingPrice = $this->totalWagerValue * ($this->sellingPercentage / 100);
        return ($value > $gtSellingPrice);
    }

    /**
     * Get the validation error message.
     *
     * @return string
     */
    public function message()
    {
        return 'The :attribute must be greater than total_wager_value * (selling_percentage / 100).';
    }
}
