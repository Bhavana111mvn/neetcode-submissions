func kClosest(points [][]int, k int) [][]int {
   distances := make([][]int,0)
   finalarr := make([][]int,0)
   for ind, point := range points {
        distance := (point[0]*point[0] + point[1]*point[1])
        tmp := make([]int,2)
        tmp[0] = ind
        tmp[1]= int(distance)
        distances = append(distances, tmp)
   }
   sort.Slice(distances, func(i, j int) bool {
    if distances[i][1] == distances[j][1] {
        return distances[i][0] < distances[j][0] // lower I first if distance equal
    }
    return distances[i][1] < distances[j][1]
   })
   for ind:=0; ind<len(distances); ind++ {
      if len(finalarr) == k {
         break
      }
      finalarr = append(finalarr, points[distances[ind][0]])
   }
   return finalarr
}
