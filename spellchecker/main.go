package main

import (
	"fmt"
	"regexp"
)


const EXAMPLE_CODE = `
import React from "react";
import {Button} from "@design/button";

interface IProps {
    onClick?: (event: React.MouseEvent) => void;
    variant?: "tile";
}

export const AskForPriceButton: React.FC<IProps> = (props) => {
    const onClick = (event: React.MouseEvent) => {
        event.preventDefault();
        event.stopPropagation();
        props.onClick && props.onClick(event);
    };

    return (
        <Button onClick={onClick} variant="highlight_primary">
            <span data-testid="button-ask-for-price">{props.children}</span>
        </Button>
    );
};


export const onAskForPriceButton: React.FC<IProps> = (props) => {
    const onClick = (event: React.MouseEvent) => {
        event.preventDefault();
        event.stopPropagation();
        props.onClick && props.onClick(event);
    };

    return (
        <Button onClick={onClick} variant="highlight_primary">
            <span data-testid="button-ask-for-price">{props.children}</span>
        </Button>
    );
};


import React from "react";
import {Button} from "@design/button";

interface IProps {
    onClick?: (event: React.MouseEvent) => void;
    variant?: "tile";
}

export const AskForPriceButton: React.FC<IProps> = (props) => {
    const onClick = (event: React.MouseEvent) => {
        event.preventDefault();
        event.stopPropagation();
        props.onClick && props.onClick(event);
    };

    return (
        <Button onClick={onClick} variant="highlight_primary">
            <span data-testid="button-ask-for-price">{props.children}</span>
        </Button>
    );
};
`;

func main() {
	fmt.Println(findVariableNames(EXAMPLE_CODE))

    for _, value := range findVariableNames(EXAMPLE_CODE) {
        fmt.Printf("# %v \n", convertVariableNameIntoWorlds(value))
    }
}

func findVariableNames(code string) []string {
	r, _ := regexp.Compile("(const) (([^ :]+)*)")
	var variableNames []string
	regexResult := r.FindAllStringSubmatch(code, -1)

	for _ , value := range regexResult {
		variableNames = append(variableNames, value[2])
	}

	return variableNames
}

func convertVariableNameIntoWorlds(variableName string) []string{
    re := regexp.MustCompile("[A-Z]")
    uppercaseLettsIndexes := re.FindAllStringIndex(variableName, -1)
	var results []string

    for i, _ := range uppercaseLettsIndexes {

        fmt.Printf("uppercaseLetterIndexes for %v, are: %v \n",variableName, uppercaseLettsIndexes)

       var wordStartIndex int
       if(i == 0) {
        wordStartIndex = 0
       } else if(i == len(uppercaseLettsIndexes) - 1){
        wordStartIndex = uppercaseLettsIndexes[i][0]
       } else {
        wordStartIndex = uppercaseLettsIndexes[i - 1][0]
       }

      var wordEndIndex int

      if(i == len(uppercaseLettsIndexes) - 1) {
        fmt.Printf("index is%v, %v \n", i, len(uppercaseLettsIndexes) -1)
        wordEndIndex = len(variableName)
      } else {
         wordEndIndex = uppercaseLettsIndexes[i][0]
      }

        

       word := variableName[wordStartIndex:wordEndIndex]
       results = append(results, word)
    }

    return results
}
