package terrarium

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

type NeuralNetworkConfig struct {
	inputNeurons	int
	outputNeurons	int
	hiddenNeurons	int
	epochs		int
	learningRate	float64
}

type NeuralNetwork struct {
	config		NeuralNetworkConfig
	weightsHidden	*mat.Dense
	biasesHidden	*mat.Dense
	weightsOut	*mat.Dense
	biasesOut	*mat.Dense
}

func NewNeuralNetwork(config NeuralNetworkConfig) *NeuralNetwork {

	nn := &NeuralNetwork{ config: config }

	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	
	wHidden := mat.NewDense(nn.config.inputNeurons, nn.config.hiddenNeurons, nil)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons, nil)
	wOut := mat.NewDense(nn.config.hiddenNeurons, nn.config.outputNeurons, nil)
	bOut := mat.NewDense(1, nn.config.outputNeurons, nil)
	
	wHiddenRaw := wHidden.RawMatrix().Data
	bHiddenRaw := bHidden.RawMatrix().Data
	wOutRaw := wOut.RawMatrix().Data
	bOutRaw := bOut.RawMatrix().Data
	
	for _, param := range [][]float64{
	    wHiddenRaw,
	    bHiddenRaw,
	    wOutRaw,
	    bOutRaw,
	} {
	    for i := range param {
	        param[i] = randGen.Float64()
	    }
	}
	
	nn.weightsHidden = wHidden
	nn.biasesHidden = bHidden
	nn.weightsOut = wOut
	nn.biasesOut = bOut
	return nn
}

func (nn *NeuralNetwork) Think(x *mat.Dense) (*mat.Dense, error) {
	if nn.weightsHidden == nil || nn.weightsOut == nil {
		return nil, errors.New("the supplied weights are empty")
	}
	if nn.biasesHidden == nil || nn.biasesOut == nil {
	    return nil, errors.New("the supplied biases are empty")
	}
	
	// Define the output of the neural network.
	output := new(mat.Dense)
	
	// Complete the feed forward process.
	hiddenLayerInput := new(mat.Dense)
	hiddenLayerInput.Mul(x, nn.weightsHidden)
	addBHidden := func(_, col int, v float64) float64 { return v + nn.biasesHidden.At(0, col) }
	hiddenLayerInput.Apply(addBHidden, hiddenLayerInput)
	
	hiddenLayerActivations := new(mat.Dense)
	applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
	hiddenLayerActivations.Apply(applySigmoid, hiddenLayerInput)
	
	outputLayerInput := new(mat.Dense)
	outputLayerInput.Mul(hiddenLayerActivations, nn.weightsOut)
	addBOut := func(_, col int, v float64) float64 { return v + nn.biasesOut.At(0, col) }
	outputLayerInput.Apply(addBOut, outputLayerInput)
	output.Apply(applySigmoid, outputLayerInput)
	
	return output, nil
}

func sumAlongAxis(axis int, m *mat.Dense) (*mat.Dense, error) {

        numRows, numCols := m.Dims()

        var output *mat.Dense

        switch axis {
        case 0:
                data := make([]float64, numCols)
                for i := 0; i < numCols; i++ {
                        col := mat.Col(nil, i, m)
                        data[i] = floats.Sum(col)
                }
                output = mat.NewDense(1, numCols, data)
        case 1:
                data := make([]float64, numRows)
                for i := 0; i < numRows; i++ {
                        row := mat.Row(nil, i, m)
                        data[i] = floats.Sum(row)
                }
                output = mat.NewDense(numRows, 1, data)
        default:
                return nil, errors.New("invalid axis, must be 0 or 1")
        }

        return output, nil
}

func (nn *NeuralNetwork) backpropagate(x, y, wHidden, bHidden, wOut, bOut, output *mat.Dense) error {
	for i:= 0; i < nn.config.epochs; i++ {
		hiddenLayerInput := new(mat.Dense)
		hiddenLayerInput.Mul(x, wHidden)
		addBiasesHidden := func(_, col int, v float64) float64 {
			return v + bHidden.At(0, col)
		}
		hiddenLayerInput.Apply(addBiasesHidden, hiddenLayerInput)

		hiddenLayerActivations := new(mat.Dense)
		applySigmoid := func(_, _ int, v float64) float64 {
			return sigmoid(v)
		}
		hiddenLayerActivations.Apply(applySigmoid, hiddenLayerInput)

		outputLayerInput := new(mat.Dense)
		outputLayerInput.Mul(hiddenLayerActivations, wOut)
		addBiasesOut := func(_, col int, v float64) float64 {
			return v + bOut.At(0, col)
		}
		outputLayerInput.Apply(addBiasesOut, outputLayerInput)

		networkError := new(mat.Dense)
		networkError.Sub(y, output)
		slopeOutputLayer := new(mat.Dense)
		applySigmoidPrime := func(_, _ int, v float64) float64 { return sigmoidPrime(v) }
		slopeOutputLayer.Apply(applySigmoidPrime, output)
		slopeHiddenLayer := new(mat.Dense)
		slopeHiddenLayer.Apply(applySigmoidPrime, hiddenLayerActivations)
		
		dOutput := new(mat.Dense)
		dOutput.MulElem(networkError, slopeOutputLayer)
		errorAtHiddenLayer := new(mat.Dense)
		errorAtHiddenLayer.Mul(dOutput, wOut.T())
		
		dHiddenLayer := new(mat.Dense)
		dHiddenLayer.MulElem(errorAtHiddenLayer, slopeHiddenLayer)
		
		wOutAdj := new(mat.Dense)
		wOutAdj.Mul(hiddenLayerActivations.T(), dOutput)
		wOutAdj.Scale(nn.config.learningRate, wOutAdj)
		wOut.Add(wOut, wOutAdj)
		
		bOutAdj, err := sumAlongAxis(0, dOutput)
		if err != nil {
		    return err
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOut.Add(bOut, bOutAdj)
		
		wHiddenAdj := new(mat.Dense)
		wHiddenAdj.Mul(x.T(), dHiddenLayer)
		wHiddenAdj.Scale(nn.config.learningRate, wHiddenAdj)
		wHidden.Add(wHidden, wHiddenAdj)
		
		bHiddenAdj, err := sumAlongAxis(0, dHiddenLayer)
		if err != nil {
		    return err
		}
		bHiddenAdj.Scale(nn.config.learningRate, bHiddenAdj)
		bHidden.Add(bHidden, bHiddenAdj)
 
	}

	return nil
}

func (nn *NeuralNetwork) Improve(x, y, output, hiddenLayerActivations *mat.Dense) error {
	networkError := new(mat.Dense)
	networkError.Sub(y, output)
	slopeOutputLayer := new(mat.Dense)
	applySigmoidPrime := func(_, _ int, v float64) float64 { return sigmoidPrime(v) }
	slopeOutputLayer.Apply(applySigmoidPrime, output)
	slopeHiddenLayer := new(mat.Dense)
	slopeHiddenLayer.Apply(applySigmoidPrime, hiddenLayerActivations)
	
	dOutput := new(mat.Dense)
	dOutput.MulElem(networkError, slopeOutputLayer)
	errorAtHiddenLayer := new(mat.Dense)
	errorAtHiddenLayer.Mul(dOutput, nn.weightsOut.T())
	
	dHiddenLayer := new(mat.Dense)
	dHiddenLayer.MulElem(errorAtHiddenLayer, slopeHiddenLayer)
	
	wOutAdj := new(mat.Dense)
	wOutAdj.Mul(hiddenLayerActivations.T(), dOutput)
	wOutAdj.Scale(nn.config.learningRate, wOutAdj)
	nn.weightsOut.Add(nn.weightsOut, wOutAdj)
	
	bOutAdj, err := sumAlongAxis(0, dOutput)
	if err != nil {
	    return err
	}
	bOutAdj.Scale(nn.config.learningRate, bOutAdj)
	nn.biasesOut.Add(nn.biasesOut, bOutAdj)
	
	wHiddenAdj := new(mat.Dense)
	wHiddenAdj.Mul(x.T(), dHiddenLayer)
	wHiddenAdj.Scale(nn.config.learningRate, wHiddenAdj)
	nn.weightsHidden.Add(nn.weightsHidden, wHiddenAdj)
	
	bHiddenAdj, err := sumAlongAxis(0, dHiddenLayer)
	if err != nil {
	    return err
	}
	bHiddenAdj.Scale(nn.config.learningRate, bHiddenAdj)
	nn.biasesHidden.Add(nn.biasesHidden, bHiddenAdj)

	return nil
}

func (nn *NeuralNetwork) Learn(x, y *mat.Dense) error {
	output := new(mat.Dense)
	
	if err := nn.backpropagate(x, y, nn.weightsHidden, nn.biasesHidden, nn.weightsOut, nn.biasesOut, output); err != nil {
	    return err
	}
	return nil
}

func sigmoid(x float64) float64 {
	return 1.0/(1.0 + math.Exp(-x))
}

func sigmoidPrime(x float64) float64 {
	return sigmoid(x) * (1.0 - sigmoid(x))
}
